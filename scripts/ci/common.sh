#!/bin/bash

repo_name() {
    echo "registry"
}

aws_region() {
    echo "us-west-2"
}

# Posts a message to Slack. Requires a valid access token is available in $SLACK_ACCESS_TOKEN.
# Usage: post_to_slack <channel> <message>
post_to_slack() {
    local channel=$1
    local message=$2

    local escaped=$(echo ${message} | sed 's/"/\"/g' | sed "s/'/\'/g" )
    local json="{\"channel\": \"#${channel}\", \"text\": \"${escaped}\", \"as_user\": true}"

    curl -s \
         -X POST \
         -H "Content-type: application/json" \
         -H "Authorization: Bearer ${SLACK_ACCESS_TOKEN}" \
         -d  "${json}" \
         https://slack.com/api/chat.postMessage > /dev/null
}

# Posts a comment to a GitHub PR. Requires a GitHub token is available in $GITHUB_TOKEN.
# Usage: post_github_pr_comment "Hi!" "https://api.github.com/repos/<org>/<repo>/issues/<pr-number>/comments"
post_github_pr_comment() {
    local pr_comment=$1
    local pr_comment_api_url=$2
    local pr_comment_body=$(printf '{ "body": "%s" }' "$pr_comment")

    curl -s \
         -X POST \
         -H "Authorization: token ${GITHUB_TOKEN}" \
         -d "$pr_comment_body" \
         $pr_comment_api_url > /dev/null
}

# The logins the preview comment can legitimately be authored by. The pull_request build and
# the /preview command both run with secrets.GITHUB_TOKEN, so in practice it's always
# github-actions[bot]; pulumi-bot is here so a run authenticated with the bot PAT adopts the
# same comment instead of starting a second one. Matching on the author as well as the marker
# also keeps a contributor from redirecting the pinned comment by pasting the marker into a
# comment of their own.
pr_comment_author_logins() {
    echo '["github-actions[bot]", "pulumi-bot"]'
}

# Creates a PR comment, or updates the existing one carrying the same marker. This is what
# makes the preview comment "pinned": one comment per PR, rewritten on every build, rather
# than a fresh comment per commit.
# Usage: upsert_github_pr_comment <marker> <body> <comments-api-url> <repo-api-url>
#   <comments-api-url> is .pull_request._links.comments.href from the event payload.
#   <repo-api-url>     is .pull_request.base.repo.url from the event payload.
upsert_github_pr_comment() {
    local marker=$1
    local body=$2
    local pr_comment_api_url=$3
    local repo_api_url=$4

    local payload
    payload=$(jq -n --arg body "$body" '{"body": $body}')

    # Find the comment to update. The list has to be paginated: GitHub returns issue comments
    # oldest-first, 30 per page by default, so an unpaginated search only ever sees the 30
    # oldest. The pinned comment is usually among them -- but not on a PR that collected 30+
    # comments before its first successful preview build (fact sheets, /check runs, review
    # chatter, or earlier builds that failed before they got this far). Once it falls off page
    # one it is never found again, and every later build appends another duplicate.
    local max_comment_pages=10
    local existing_comment_id=""
    local page comments_json page_count
    for page in $(seq 1 "$max_comment_pages"); do
        comments_json=$(curl -s \
            -H "Authorization: token ${GITHUB_TOKEN}" \
            "${pr_comment_api_url}?per_page=100&page=${page}")

        local found
        found=$(echo "$comments_json" | jq -r \
            --arg marker "$marker" \
            --argjson logins "$(pr_comment_author_logins)" \
            '[.[] | select((.body // "") | contains($marker))
                  | select(.user.login as $l | $logins | index($l))] | last | .id // empty' \
            2>/dev/null)
        if [[ -n "$found" ]]; then
            existing_comment_id="$found"
        fi

        # Stop on the last (short) page, or if the response wasn't an array — a transient API
        # error yields 0 here and breaks cleanly rather than looping ten times.
        page_count="$(echo "$comments_json" | jq -r 'if type == "array" then length else 0 end' 2>/dev/null)"
        if [[ "${page_count:-0}" -lt 100 ]]; then
            break
        fi
    done

    if [[ -n "$existing_comment_id" ]]; then
        curl -s \
            -X PATCH \
            -H "Authorization: token ${GITHUB_TOKEN}" \
            -H "Content-Type: application/json" \
            -d "$payload" \
            "${repo_api_url}/issues/comments/${existing_comment_id}" > /dev/null
    else
        curl -s \
            -X POST \
            -H "Authorization: token ${GITHUB_TOKEN}" \
            -H "Content-Type: application/json" \
            -d "$payload" \
            "$pr_comment_api_url" > /dev/null
    fi
}

# The Hugo content root. Hugo mounts themes/default/content as the site's content directory
# (see the module import in config/_default/config.yml), so an on-disk path under here maps
# directly onto a site path -- including the leading /registry/ segment.
content_root() {
    echo "themes/default/content/"
}

# Converts a Hugo content file path to its published, root-relative URL, following the same
# rules the site uses: strip the content root, strip the ".md" extension, collapse section
# landing pages ("/_index") and leaf-bundle pages ("/index") to their directory, and guarantee
# a leading and trailing slash.
# Usage:
#   content_path_to_url themes/default/content/registry/packages/aws/_index.md
#     # => /registry/packages/aws/
#   content_path_to_url themes/default/content/registry/packages/aws/installation-configuration.md
#     # => /registry/packages/aws/installation-configuration/
content_path_to_url() {
    local path=$1
    local root
    root="$(content_root)"

    path="${path#"$root"}"           # Strip the leading content root.
    path="${path%.md}"               # Strip the trailing .md extension.

    if [[ "$path" == "_index" || "$path" == "index" ]]; then
        path=""                      # Site root.
    else
        path="${path%/_index}"       # Collapse section landing pages (_index.md).
        path="${path%/index}"        # Collapse leaf-bundle pages (index.md).
    fi

    echo "/${path}/" | sed 's#//*#/#g'
}

# Maps changed repository paths (read from stdin, one per line) to the site URLs they render,
# de-duplicated and in first-seen order. Anything that isn't a page source is dropped.
#
# Two rules, because the registry has two kinds of page source:
#   - Content markdown under themes/default/content/.
#   - Package metadata under themes/default/data/registry/packages/<pkg>.yaml, which is what
#     actually renders /registry/packages/<pkg>/. This rule is the important one here: the
#     generated api-docs/ content is gitignored and never appears in a PR diff, so without it
#     the changed-page list would be empty on most registry PRs.
#
# The <pkg> slug is the YAML's basename. That matches the file's `name:` field for every
# package in the registry today, and callers existence-check the result anyway, so a
# hypothetical divergence drops the link rather than emitting a dead one.
changed_paths_to_urls() {
    local path url root
    local -A seen=()
    root="$(content_root)"

    while IFS= read -r path; do
        [[ -z "$path" ]] && continue

        url=""
        case "$path" in
            "$root"*.md)
                url="$(content_path_to_url "$path")"
                ;;
            themes/default/data/registry/packages/*.yaml)
                url="/registry/packages/$(basename "$path" .yaml)/"
                ;;
        esac

        [[ -z "$url" ]] && continue
        [[ -n "${seen[$url]:-}" ]] && continue
        seen[$url]=1
        echo "$url"
    done
}

# Returns the Git SHA of the HEAD commit. For pull requests, we take this from GitHub event metadata, since in that case, the HEAD commit will contain the SHA of the merge commit with the base branch.
git_sha() {
    if [[ "$GITHUB_EVENT_NAME" == "pull_request" && ! -z "$GITHUB_EVENT_PATH" ]]; then
        echo "$(cat "$GITHUB_EVENT_PATH" | jq -r ".pull_request.head.sha")"
    else
        echo "$(git rev-parse HEAD)"
    fi
}

# Returns the shortened version of either the GITHUB_SHA, if present, or that of the most
# recent commit.
git_sha_short() {
    echo "$(git_sha)" | cut -c1-8
}

# current_time_in_ms returns the epoch time in milliseconds.
current_time_in_ms() {
    echo "$(node -e 'console.log(Date.now())')"
}

origin_bucket_prefix() {
    echo "$(repo_name)-${DEPLOYMENT_ENVIRONMENT}-origin"
}

# Returns the name of the metadata file we expect to exist locally before running Pulumi.
origin_bucket_metadata_filepath() {
    echo "./origin-bucket-metadata.json"
}

# build_identifier returns a string that is used to identify the current build for naming
# S3 buckets and asset bundles.
build_identifier() {
    local identifier

    # For CI builds, we use the GitHub Actions event to generate more readable identifiers.
    # - For pull_request actions, return "pr-<number>-<git-sha>"
    # - For others, return "<event-name>-<git-sha>".
    if [[ ! -z "$GITHUB_EVENT_NAME" && ! -z "$GITHUB_EVENT_PATH" ]]; then
        identifier="$GITHUB_EVENT_NAME"

        if [ "$GITHUB_EVENT_NAME" == "pull_request" ]; then
            identifier="pr-$(cat "$GITHUB_EVENT_PATH" | jq -r ".number")"
        fi

        identifier="${identifier}-$(git_sha_short)"
    else
        # For on-demand builds, if an identifier's been set, use it.
        identifier="$BUILD_IDENTIFIER"

        # Otherwise, just use the current Git SHA.
        if [ -z "$BUILD_IDENTIFIER" ]; then
            identifier="$(git_sha_short)"
        fi
    fi

    echo "$identifier"
}

# Get the AWS SSM Parameter Store key for the specified commit SHA. Used for mapping a
# commit to a previously created bucket.
ssm_parameter_key_for_commit() {
    echo "/$(repo_name)/commits/$1/bucket"
}

# Get the S3 bucket associated with a specific commit.
get_bucket_for_commit() {
    aws ssm get-parameter \
        --name "$(ssm_parameter_key_for_commit $1)" \
        --query Parameter.Value \
        --region us-west-2 \
        --output text || echo ""
}

aws_owner_tag() {
    echo "Key=Owner,Value=pulumi-hugo"
}

# Set the S3 bucket associated with a specific commit.
set_bucket_for_commit() {
    aws ssm put-parameter \
        --name "$(ssm_parameter_key_for_commit $1)" \
        --value "$2" \
        --type String \
        --region $3 \
        --overwrite
}

# Remove the parameter key associated with a specific commit.
remove_param_for_commit() {
    # SSM errors when the parameter doesn't exist (e.g. somebody manually deleted it), so we check and ignore that error.
    cmd_status=0
    output=$(aws ssm delete-parameter --name "$(ssm_parameter_key_for_commit $1)" --region $2 2>&1) || cmd_status=$?
    if [[ $cmd_status -eq 0 ]] || [[ $output == *"ParameterNotFound"* ]]; then
        return 0
    fi
    echo "$output" >&2
    return 1
}

# List the 100 most recent bucket in the current account, sorted descendingly by
# CreationDate, matching the prefix we use to name website buckets. Supports an optional
# suffix to filter by (e.g., "pr" or "push").
get_recent_buckets() {
    aws s3api list-buckets \
        --query "reverse(sort_by(Buckets,&CreationDate))[:100].{id:Name,date:CreationDate}|[?starts_with(id,'$(origin_bucket_prefix)-${1}')]" \
        --output json | jq -r '.[].id'
}

# Retry the given command some number of times, with a delay of some number of seconds between calls.
# Usage: retry some_command <retry-count> <delay-in-seconds>
retry() {
    local n=1
    local max=$2
    local delay=$3
    while true; do
    "$@" && break || {
        if [[ $n -lt $max ]]; then
            ((n++))
            echo "Command failed. Attempt $n/$max:"
            sleep $delay;
        else
            echo "The command has failed after $n attempts." >&2
            return 1
        fi
    }
    done
}


# Provides a time-stamped log entry.
log() {
    echo "[$(date "+%Y.%m.%d:%H.%M.%S")]" "$@"
}

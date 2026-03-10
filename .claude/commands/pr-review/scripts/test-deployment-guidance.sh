#!/bin/bash
# Test Deployment Guidance Script for Pulumi Registry
# Usage: ./test-deployment-guidance.sh <PR_NUMBER>
#
# Outputs:
# JSON object with deployment URL, status, and categorized file data

set -euo pipefail

if [ -z "$1" ]; then
  echo "Error: PR number required" >&2
  echo "Usage: $0 <PR_NUMBER>" >&2
  exit 1
fi

PR_NUMBER="$1"

# Fetch pulumi-bot comment and extract preview deployment URL
# Registry pattern: registry--origin-pr-* S3 bucket
DEPLOYMENT_URL=$(gh api repos/pulumi/registry/issues/"$PR_NUMBER"/comments \
  --jq '.[] | select(.user.login == "pulumi-bot") | .body' 2>/dev/null | \
  grep -oP 'https?://[a-zA-Z0-9._-]*registry[a-zA-Z0-9._-]*origin-pr-[a-zA-Z0-9._-]*\.s3-website[a-zA-Z0-9._-]*\.amazonaws\.com' | \
  tail -1 || echo "")

# Check for automation label (metadata-only PRs)
PR_LABELS=$(gh pr view "$PR_NUMBER" --json labels --jq '.labels[].name' 2>/dev/null || echo "")
IS_AUTOMATION="false"
if echo "$PR_LABELS" | grep -q "automation/tfgen-provider-docs"; then
  IS_AUTOMATION="true"
fi

# Fetch PR data including files
PR_DATA=$(gh pr view "$PR_NUMBER" --json files,additions,deletions)

# Get changed files
CHANGED_FILES=$(echo "$PR_DATA" | jq -r '.files[].path')

# Categorize files
PACKAGE_YAML=()
GO_TOOLS=()
CI_WORKFLOWS=()
BUILD_SCRIPTS=()
INFRASTRUCTURE=()
THEME_LAYOUT=()
COMMUNITY_PACKAGES=()
OTHER=()

while IFS= read -r file; do
  [ -z "$file" ] && continue
  case "$file" in
    themes/default/data/registry/packages/*.yaml)
      PACKAGE_YAML+=("$file") ;;
    tools/resourcedocsgen/*|tools/mktutorial/*)
      GO_TOOLS+=("$file") ;;
    .github/workflows/*)
      CI_WORKFLOWS+=("$file") ;;
    Makefile|scripts/*)
      BUILD_SCRIPTS+=("$file") ;;
    infrastructure/*)
      INFRASTRUCTURE+=("$file") ;;
    themes/default/layouts/*|themes/default/theme/*)
      THEME_LAYOUT+=("$file") ;;
    community-packages/*)
      COMMUNITY_PACKAGES+=("$file") ;;
    *)
      OTHER+=("$file") ;;
  esac
done <<< "$CHANGED_FILES"

# Helper to create JSON array from bash array
json_array() {
  local arr=("$@")
  if [ ${#arr[@]} -eq 0 ]; then
    echo "[]"
    return
  fi
  printf '%s\n' "${arr[@]}" | jq -R . | jq -s .
}

# Build output JSON
jq -n \
  --arg prNumber "$PR_NUMBER" \
  --arg deploymentUrl "$DEPLOYMENT_URL" \
  --arg isAutomation "$IS_AUTOMATION" \
  --argjson packageYaml "$(json_array "${PACKAGE_YAML[@]}")" \
  --argjson goTools "$(json_array "${GO_TOOLS[@]}")" \
  --argjson ciWorkflows "$(json_array "${CI_WORKFLOWS[@]}")" \
  --argjson buildScripts "$(json_array "${BUILD_SCRIPTS[@]}")" \
  --argjson infrastructure "$(json_array "${INFRASTRUCTURE[@]}")" \
  --argjson themeLayout "$(json_array "${THEME_LAYOUT[@]}")" \
  --argjson communityPackages "$(json_array "${COMMUNITY_PACKAGES[@]}")" \
  --argjson other "$(json_array "${OTHER[@]}")" \
  '{
    prNumber: ($prNumber | tonumber),
    deploymentUrl: (if $deploymentUrl == "" then null else $deploymentUrl end),
    deploymentStatus: (if $deploymentUrl == "" then "pending" else "ready" end),
    isAutomationPR: ($isAutomation == "true"),
    fileCategories: {
      packageYaml: $packageYaml,
      goTools: $goTools,
      ciWorkflows: $ciWorkflows,
      buildScripts: $buildScripts,
      infrastructure: $infrastructure,
      themeLayout: $themeLayout,
      communityPackages: $communityPackages,
      other: $other
    }
  }'

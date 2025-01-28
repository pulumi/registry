.PHONY: clean
clean:
	./scripts/clean.sh

.PHONY: ensure
ensure:
	./scripts/ensure.sh

.PHONY: lint
lint: lint-go lint-markdown
	yarn run lint

.PHONY: lint-markdown
lint-markdown:
	./scripts/lint/lint-markdown.js

.PHONY: lint-go
lint-go: lint-resourcedocsgen lint-mktutorial

.PHONY: lint-resourcedocsgen
lint-resourcedocsgen:
	cd tools/resourcedocsgen/ && golangci-lint run --config ../../.golangci.yml --path-prefix tools/resourcedocsgen/

.PHONY: lint-mktutorial
lint-mktutorial:
	cd tools/mktutorial/ && golangci-lint run --config ../../.golangci.yml --path-prefix tools/mktutorial/

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build: build-assets
	node scripts/apply-fixes.js
	hugo | grep -v -e 'WARN .* REF_NOT_FOUND'

HUGO_SERVE := hugo serve --buildDrafts --buildFuture --ignoreVendorPaths="github.com/pulumi/registry/**/*" | grep -v -e 'WARN .* REF_NOT_FOUND'

# bin/resourcedocsgen is considered to depend on all files in the tools/resourcedocsgen directory.
#
# Caveat: This does not account for `replace` directives in `go.mod`.
bin/resourcedocsgen: $(shell find tools/resourcedocsgen | sed 's@ @\\ @g')
	mkdir -p bin
	go build -C tools/resourcedocsgen -o ../../bin ./...

# Generate the API docs for `content/registry/packages/<pkg>`, where <pkg> is the name of
# the package to be handed off to resourcedocsgen.
content/registry/packages/%: bin/resourcedocsgen
	bin/resourcedocsgen docs registry \
		--baseDocsOutDir ./content/registry/packages \
		--basePackageTreeJSONOutDir ./static/registry/packages/navs \
		$*

# Run hugo locally, ignoring REF_NOT_FOUND warnings
.PHONY: serve
serve:
	${HUGO_SERVE}

.PHONY: ci-pull-request
ci-pull-request: ensure
	./scripts/ci/pull-request.sh

.PHONY: ci-pull-request-closed
ci-pull-request-closed:
	./scripts/ci/pull-request-closed.sh

.PHONY: ci-scheduled
ci-scheduled:
	./scripts/ci/scheduled.sh

.PHONY: ci_push
ci_push:: ensure
	./scripts/ci/push.sh

.PHONY: serve-assets
serve-assets:
	yarn --cwd ./themes/default/theme run start

.PHONY: serve-all
serve-all:
	./node_modules/.bin/concurrently --kill-others -r "${HUGO_SERVE}" "yarn --cwd ./themes/default/theme run start"

.PHONY: build-assets
build-assets:
	yarn --cwd ./themes/default/theme run build

.PHONY: ci_bucket_cleanup
ci_bucket_cleanup:
	./scripts/ci/bucket-cleanup.sh

.PHONY: check_links
check_links: ensure
	./scripts/link-checker/check-links.sh "https://www.pulumi.com/registry"

.PHONY: run-browser-tests
run-browser-tests: ensure
	./scripts/run-api-docs-tests.sh

.PHONY: clean
clean:
	./scripts/clean.sh

.PHONY: ensure
ensure:
	./scripts/ensure.sh

.PHONY: lint
lint: lint-go lint-markdown

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
build:
	$(MAKE) build-assets
	./scripts/build.sh

HUGO_SERVE := hugo serve --buildDrafts --buildFuture --ignoreVendorPaths="github.com/pulumi/registry/**/*" | grep -v -e 'WARN .* REF_NOT_FOUND'

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
ci_push::
	$(MAKE) ensure
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
check_links:
	$(MAKE) ensure
	./scripts/link-checker/check-links.sh "https://www.pulumi.com/registry"

.PHONY: run-browser-tests
run-browser-tests:
	$(MAKE) ensure
	./scripts/run-api-docs-tests.sh

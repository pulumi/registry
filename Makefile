.PHONY: clean
clean:
	./scripts/clean.sh

.PHONY: ensure
ensure:
	./scripts/ensure.sh

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: lint-go
lint-go: lint-resourcedocsgen

.PHONY: lint-resourcedocsgen
lint-resourcedocsgen:
	cd tools/resourcedocsgen/ && golangci-lint run --config ../../.golangci.yml --path-prefix tools/resourcedocsgen/

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build:
	$(MAKE) build-assets
	./scripts/build.sh

.PHONY: serve
serve:
	./scripts/serve.sh

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
	./node_modules/.bin/concurrently --kill-others -r "./scripts/serve.sh" "yarn --cwd ./themes/default/theme run start"

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

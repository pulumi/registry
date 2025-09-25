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
	cd tools/resourcedocsgen/ && golangci-lint run --config ../../.golangci.yml

.PHONY: lint-mktutorial
lint-mktutorial:
	cd tools/mktutorial/ && golangci-lint run --config ../../.golangci.yml

.PHONY: test
test:
	cd ./tools/resourcedocsgen && go test ./...

.PHONY: build
build: build-assets
	node scripts/apply-fixes.js
	hugo | grep -v -e 'WARN .* REF_NOT_FOUND'

HUGO_SERVE := hugo serve --buildDrafts --buildFuture --ignoreVendorPaths="github.com/pulumi/registry/**/*" | grep -v -e 'WARN .* REF_NOT_FOUND'

HELPMAKEGO_VERSION := v0.1.0
HELPMAKEGO := bin/${HELPMAKEGO_VERSION}/helpmakego

# Ensure that `helpmakego` is installed at ${HELPMAKEGO} before it is used to resolve targets.
#
# This has the side effect of ensuring that the `bin` directory is present.
_ := $(shell if ! [ -x ${HELPMAKEGO} ]; then \
	GOBIN=$(shell pwd)/bin/${HELPMAKEGO_VERSION} go install github.com/iwahbe/helpmakego@${HELPMAKEGO_VERSION}; \
	fi \
)

bin/resourcedocsgen: $(shell ${HELPMAKEGO} tools/resourcedocsgen)
	go build -C tools/resourcedocsgen -o ../../bin ./...

bin/mktutorial: $(shell ${HELPMAKEGO} tools/mktutorial)
	go build -C tools/mktutorial -o ../../bin ./...

# Generate the API docs for `content/registry/packages/<pkg>`, where <pkg> is the name of
# the package to be handed off to resourcedocsgen.
.SECONDEXPANSION: # Needed for content/registry/packages/% to have dependencies that reference %.
api-docs/%: .make/content/registry/packages/$$*/api-docs
.make/content/registry/packages/%/api-docs: bin/resourcedocsgen \
				themes/default/data/registry/packages/%.yaml \
				$$(wildcard themes/default/content/registry/packages/%/*)
	bin/resourcedocsgen docs registry \
		--baseDocsOutDir ./content/registry/packages \
		--basePackageTreeJSONOutDir ./static/registry/packages/navs \
		$*
	@mkdir -p "$(@D)"
	@touch $@

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
build-assets: ensure
	yarn --cwd ./themes/default/theme run build

.PHONY: ci_bucket_cleanup
ci_bucket_cleanup:
	./scripts/ci/bucket-cleanup.sh

.PHONY: check_links
check_links: ensure
	yarn run check-links

.PHONY: test_provider_api_docs
test_provider_api_docs: ensure build-assets bin/resourcedocsgen
	./scripts/ci/run-provider-tests.sh

.PHONY: run-browser-tests
run-browser-tests: ensure
	./scripts/run-api-docs-tests.sh

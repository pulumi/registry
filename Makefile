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
test: test-infra
	cd ./tools/resourcedocsgen && go test ./...

.PHONY: test-infra
test-infra:
	cd infrastructure && node --test redirects.test.js

.PHONY: build
build: build-assets
	node scripts/apply-fixes.js
	hugo 2>&1 | grep -v -e 'WARN .* REF_NOT_FOUND'
	node scripts/split-sitemap.js

HUGO_SERVE := hugo serve --buildDrafts --buildFuture --ignoreVendorPaths="github.com/pulumi/registry/**/*" 2>&1 | grep -v -e 'WARN .* REF_NOT_FOUND'

HELPMAKEGO_VERSION := v0.1.0
HELPMAKEGO := bin/${HELPMAKEGO_VERSION}/helpmakego

# Ensure that `helpmakego` is installed at ${HELPMAKEGO} before it is used to resolve targets.
#
# This has the side effect of ensuring that the `bin` directory is present.
_ := $(shell if ! [ -x ${HELPMAKEGO} ]; then \
	GOBIN=$(shell pwd)/bin/${HELPMAKEGO_VERSION} go install github.com/iwahbe/helpmakego@${HELPMAKEGO_VERSION}; \
	fi \
)

DOCSGEN_SRC_HASH := $(shell find tools/resourcedocsgen -name '*.go' -o -name '*.tmpl' -o -name 'go.sum' | sort | xargs sha256sum | sha256sum | cut -d' ' -f1)

bin/resourcedocsgen: $(shell ${HELPMAKEGO} tools/resourcedocsgen)
	go build -C tools/resourcedocsgen -ldflags "-X github.com/pulumi/registry/tools/resourcedocsgen/cmd/docs.sourceHash=$(DOCSGEN_SRC_HASH)" -o ../../bin ./...

bin/mktutorial: $(shell ${HELPMAKEGO} tools/mktutorial)
	go build -C tools/mktutorial -o ../../bin ./...

# Generate API docs for all packages, then versioned docs for blessed packages.
# Set SKIP_VERSIONED_DOCS=1 to skip versioned docs generation (e.g. when running in a parallel CI job).
.PHONY: api-docs
api-docs: bin/resourcedocsgen
	bin/resourcedocsgen docs registry \
		--baseDocsOutDir ./themes/default/content/registry/packages \
		--basePackageTreeJSONOutDir ./themes/default/static/registry/packages/navs \
		--baseSchemasOutDir ./themes/default/static/registry/packages \
		--baseLLMDocsOutDir ./llm-docs-out/registry/packages \
		--logtostderr
	$(if $(SKIP_VERSIONED_DOCS),,./scripts/generate-versioned-docs.sh)

# Generate the API docs for `content/registry/packages/<pkg>`, where <pkg> is the name of
# the package to be handed off to resourcedocsgen.
.SECONDEXPANSION: # Needed for content/registry/packages/% to have dependencies that reference %.
api-docs/%: .make/content/registry/packages/$$*/api-docs ;
.make/content/registry/packages/%/api-docs: bin/resourcedocsgen \
				themes/default/data/registry/packages/%.yaml \
				scripts/generate-versioned-docs.sh \
				$$(wildcard themes/default/content/registry/packages/%/*)
	@if [ -L ./content/registry/packages/$* ]; then rm ./content/registry/packages/$*; fi
	@if [ -L ./static/registry/packages/navs/$*.json ]; then rm ./static/registry/packages/navs/$*.json; fi
	bin/resourcedocsgen docs registry \
		--baseDocsOutDir ./content/registry/packages \
		--basePackageTreeJSONOutDir ./static/registry/packages/navs \
		--baseSchemasOutDir ./static/registry/packages \
		--baseLLMDocsOutDir ./llm-docs-out/registry/packages \
		$*
	CONTENT_DIR=$(CURDIR)/content/registry/packages STATIC_DIR=$(CURDIR)/static/registry/packages ./scripts/generate-versioned-docs.sh $*
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

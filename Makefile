.PHONY: clean
clean:
	./scripts/clean.sh

.PHONY: ensure
ensure:
	./scripts/ensure.sh

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: test
test:
	./scripts/test.sh

.PHONY: build
build:
	./scripts/build.sh

.PHONY: serve
serve:
	./scripts/serve.sh

.PHONY: ci-pull-request
ci-pull-request: ensure lint test
	./scripts/ci/pull-request.sh

.PHONY: ci-pull-request-closed
ci-pull-request-closed:
	./scripts/ci/pull-request-closed.sh

.PHONY: ci-scheduled
ci-scheduled:
	./scripts/ci/scheduled.sh

.PHONY: new-blog-post
new-blog-post:
	./scripts/new/blog-post.sh "$(shell bash -c 'read -p "Slug (e.g., 'my-new-post'): " slug; echo $$slug')"

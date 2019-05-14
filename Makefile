test: PKG=...
test:
	GO111MODULE=on go test -mod=vendor -cover ./$(PKG)
.PHONY: test

test.fuzz: PKG=...
test.fuzz:
	GO111MODULE=on go test -mod=vendor -cover -tags fuzz ./$(PKG)
.PHONY: test.fuzz

test.v: PKG=...
test.v:
	GO111MODULE=on go test -mod=vendor -cover -v ./$(PKG)
.PHONY: test.v

benchmark: PKG=...
benchmark: REGEX=.
benchmark:
	GO111MODULE=on go test -mod=vendor -bench $(REGEX) ./$(PKG)
.PHONY: benchmark

benchmark.full: PKG=...
benchmark.full: REGEX=.
benchmark.full: N=5
benchmark.full: FILE=./.benchmark.out
benchmark.full:
	GO111MODULE=on go test -mod=vendor \
		-bench $(REGEX) \
		-count $(N) \
		-benchmem \
		./$(PKG) > $(FILE)
	@benchstat -sort name $(FILE)
.PHONY: benchmark.full

vendor:
	GO111MODULE=on go mod vendor
.PHONY: vendor

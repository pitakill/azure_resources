test-with-clean: go-test-clean go-test-verbose
.PHONY: test-with-clean

generate: go-generate
test: go-test
test-verbose: go-test-verbose
test-clean: go-test-clean

go-generate:
	@echo "  >  Generating code"
	@cd generate && go generate

go-test:
	@echo "  >  Testing"
	go test -race . $(ARGS)

go-test-clean:
	@echo "  >  Cleaning test"
	@go clean -testcache

go-test-verbose:
	@echo "  >  Testing verbosely"
	go test -v -race . $(ARGS)

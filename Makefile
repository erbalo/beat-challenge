GO ?= go

MAIN = "main.go"
CLI_DIR = "./bin/fare-cli"

TEST_DIR_RECURSIVELY = "./..."

cli:
	$(MAKE) test
	$(GO) build -o $(CLI_DIR) $(MAIN)

test:
	$(GO) test -v $(TEST_DIR_RECURSIVELY) --parallel 10

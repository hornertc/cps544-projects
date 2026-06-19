all: test

.PHONY: test
test:
	./test.sh

# See the README.md for how to install shellcheck
.PHONY: lint
test: lint
lint:
	shellcheck src/dbfile

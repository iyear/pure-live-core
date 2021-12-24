.PHONY:build
build:
	goreleaser release --skip-publish --snapshot --rm-dist
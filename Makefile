.PHONY: test
default: test

t: test
test: lint unit-tests integration-tests acceptance-tests

ci:
	git pull -r
	make test
	git push

unit-tests:
	go test -shuffle=on --tags=unit ./...

integration-tests:
	go test -count=1 --tags=integration ./...

acceptance-tests:
	go test -count=1 --tags=acceptance ./...

generate:
	@go generate ./...

#===Linting===#
lint:
	golangci-lint run --timeout=5m

lf: lintfix
lintfix:
	golangci-lint run ./... --fix

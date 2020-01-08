.PHONY: build
build:
	@go build -o godict

.PHONY: clean
clean:
	@rm -f godict

.PHONY: test
test:
	go test -v -cover ./matcher
	go test -v -cover ./statik

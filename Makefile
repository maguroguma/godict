godict:
	@go build -o godict

.PHONY: clean
clean:
	@rm -f godict

.PHONY: test
test:
	go test -v -cover ./matcher
	go test -v -cover ./statik
	go test -v -cover ./decorator

.PHONY: install
install:
	go install

.PHONY: build

build:
	make prepare
	go build -o ./build/postal
fmt:
	go fmt
test:
	go test
prepare:
	make tidy
	make fmt
	make test
tidy:
	go mod tidy


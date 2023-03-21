build:
	@go build -o bin/ ./...

run: build
	./bin/app

test:
	@go test -v ./...

cert:
	cd cert && ./create.sh && cd ..

.PHONY: gen clean server client test cert
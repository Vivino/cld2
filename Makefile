

build: clean
	@go generate ./...

clean:
	rm -rf lib
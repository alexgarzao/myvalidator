clean:
	rm -Rf bin/

test:
	go clean -testcache
	go test -v ./...

build: clean
	go build -o bin/ ./myvalidator/...

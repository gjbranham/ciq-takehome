.PHONY: build clean test

build:
	cd ./main && go build -o ../bin/server-access

clean:
	rm -rf ./bin

test:
	go test -v ./...
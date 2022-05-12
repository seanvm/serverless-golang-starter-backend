.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/user handlers/user/main.go

test:
	go clean -testcache
	go test -v -ldflags="-s -w"  ./handlers/user/main.go  ./handlers/user/main_test.go


clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

start:
	sls offline start --noTimeout --useDocker
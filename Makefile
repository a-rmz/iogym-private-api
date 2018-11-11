build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/parseData sessions/parseData/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/parseLogin sessions/parseLogin/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/parseLogout sessions/parseLogout/main.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose

.PHONY: remove
remove:
	sls remove --verbose

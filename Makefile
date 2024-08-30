.PHONY: build clean deploy

build:
	env GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -tags lambda.norpc -o bin/bootstrap main/main.go

clean:
	rm -rf ./bin

zip:
	zip -j bin/mongodb.zip bin/bootstrap

deploy: clean build zip
	sls deploy --verbose

dev:
	go run main/main.go
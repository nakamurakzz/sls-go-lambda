build:
	GOOS=linux GOARCH=amd64 go build -o bin/main ./src

deploy:
	sls deploy

remove:
	sls remove
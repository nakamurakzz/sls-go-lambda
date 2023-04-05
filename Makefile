AWS_PROFILE ?= default

.PHONY: build deploy remove

build:
	GOOS=linux GOARCH=amd64 go build -o bin/main ./src

local:
	sls invoke local -f hello --aws-profile $(AWS_PROFILE)

deploy: build
	sls deploy --aws-profile $(AWS_PROFILE)

remove:
	sls remove --aws-profile $(AWS_PROFILE)
.PHONY: build test

gen:
	docker run -v ${PWD}:/open-dam  openapitools/openapi-generator-cli generate -i https://raw.githubusercontent.com/open-dam/open-dam-api/master/api/openapi.yaml -g go -o /open-dam --additional-properties=packageName=opendamclient

test:
	go test -cover ./...

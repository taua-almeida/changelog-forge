update-changelog:
	@go run main.go --update-changelog

generate-json:
	@go run main.go --generate-json

run-tests:
	@go test ./... -v
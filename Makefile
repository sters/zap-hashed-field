
BUILD_DIR=./build

.PHONY: dep
dep:
	dep ensure -v

.PHONY: dep-update
dep-update:
	dep ensure -v -update

.PHONY: fmt
fmt:
	@gofmt ./...

.PHONY: test
test:
	@go test -race ./...

.PHONY: cover
cover:
	@go test -race -coverpkg=./... -coverprofile=coverage.txt ./...

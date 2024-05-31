# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


no-dirty:
	git diff --exit-code


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests (sequential mode)
test:
	CGO_ENABLED=1 go test -p 1 -v -buildvcs ./...

## test/cover: run all tests and display coverage
test/cover:
	go test -v  -buildvcs -coverprofile=./test/coverage.out ./...
	go tool cover -html=./test/coverage.out


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
push: tidy audit no-dirty
	git push

.PHONY: audit confirm help no-dirty push test test/cover tidy
	
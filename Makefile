# Change these variables as necessary.
main_package_path = ./cmd/server/main.go
binary_name = main
build_path = ./
build_temporary_path = ./build/tmp

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z '$(shell git status --porcelain)'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=${build_temporary_path}/coverage.out ./...
	go tool cover -html=${build_temporary_path}/coverage.out

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	# Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=${build_temporary_path}/bin/${binary_name} ${main_package_path}

## build/linux: build the application for Linux
.PHONY: build/linux
build/linux:
    GOOS=linux GOARCH=amd64 go build -o=${build_temporary_path}/bin/linux_amd64/${binary_name} ${main_package_path}

## run: run the application
.PHONY: run
run: build
	${build_temporary_path}/bin/${binary_name}

## run/linux/live: run the application with reloading on file changes and configuration for Linux
.PHONY: run/linux/live
run/linux/live:
	air -c ${build_path}/.air.linux.toml

## run/windows/live: run the application with reloading on file changes and configuration for Windows
.PHONY: run/windows/live
run/windows/live:
	air -c ${build_path}/.air.windows.toml

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: confirm audit no-dirty
	git push

# production/deploy: deploy the application to production
# .PHONY: production/deploy
# production/deploy: confirm audit no-dirty
# 	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=${build_temporary_path}/bin/linux_amd64/${binary_name} ${main_package_path}
# 	upx -5 ${build_temporary_path}/bin/linux_amd64/${binary_name}
# 	# Include additional deployment steps here...

# ==================================================================================== #
# DATABASES
# ==================================================================================== #
## migration: create file migrations
migration:
	@migrate create -ext sql -dir scripts/database/migration -seq $(filter-out $@,$(MAKECMDGOALS))

## migration/up: run migrations
migration/up:
	@go run scripts/database/migration/migration.go up

## migration/down: rollback migration
migration/down:
	@go run scripts/database/migration/migration.go down

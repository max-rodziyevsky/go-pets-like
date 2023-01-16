BINARY_NAME = "go-pets-like"
BINARIES = "./bin"
GITHUB = "github.com/max-rodziyevsky/go-pets-like"
GIT_LOCAL_NAME = "max-rodziyevsky"
GIT_LOCAL_EMAIL = "rodziyevskydev@gmail.com"

init:
	@echo "::> Creating a module root..."
	@go mod init ${GITHUB}
	@go mod tidy
	@echo "::> Finished!"

build:
	@echo "::> Building..."
	@go build -o ${BINARIES}/${BINARY_NAME} ./
	@echo "::> Finished!"

run:
	@go build -o ${BINARIES}/${BINARY_NAME} ./
	@${BINARIES}/${BINARY_NAME}

clean:
	@echo "::> Cleaning..."
	@go clean
	@rm -rf ${BINARIES}
	@go mod tidy
	@echo "::> Finished"

local-git:
	@git config --local user.name ${GIT_LOCAL_NAME}
	@git config --local user.email ${GIT_LOCAL_EMAIL}
	@git config --local --list

.PNONY: init build run clean
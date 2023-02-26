# Build the go binaries inside the joeconvert and joectl directories and move the binaries to the bin directory
build:
	@echo "Building joeconvert"
	@cd joeconvert && go build -o ../bin/joeconvert
# @echo "Building joectl"
# @cd joectl && go build -o ../bin/joectl

# Install the binaries to the $GOPATH/bin directory
install:
	@echo "Installing joeconvert"
	@go install ./joeconvert
# @echo "Installing joectl"
# @go install ./joectl

# Run the tests
test:
	@echo "Running tests for joeconvert"
	@cd joeconvert && go test ./...
# @echo "Running tests for joectl"
# @cd joectl && go test ./...

# Run the linter
lint:
	@echo "Running linter"
	@golint ./...
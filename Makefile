# Build the program
build:
	go build -o robot .

# Run the program
run:
	go run main.go

# Run the tests
test:
	go test -race -v -failfast ./...

# Remove the binary
clean:
	rm -f robot


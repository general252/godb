

all:
	go mod tidy
	gofmt -d -l -e -w .
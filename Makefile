

all:
	go mod tidy
	goimports -w .
	gofmt -d -l -e -w .

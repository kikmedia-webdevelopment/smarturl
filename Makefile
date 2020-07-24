GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test

.PHONY: windows
windows:
	GOOS=windows GOARCH=386 go build -o mchurl.exe main.go
.PHONE: test
test:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
    $(GOCOVER) -html=coverage.out
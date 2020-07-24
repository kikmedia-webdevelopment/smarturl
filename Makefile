.PHONY: windows
windows:
	GOOS=windows GOARCH=386 go build -o mchurl.exe main.go
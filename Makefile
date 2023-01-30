main: linux

linux: init
	CGO_ENABLED=1 GOARCH=amd64 GOOS=linux  go build -o bin/gogotools_linux_amd64 main.go

windows: init
	CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build -o bin/gogotools_windows_amd64.exe main.go

darwin: init
	CGO_ENABLED=1 GOARCH=amd64 GOOS=darwin go build -o bin/gogotools_darwin_amd64.app main.go

init:
	echo "开始打包..."
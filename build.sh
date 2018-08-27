# linux amd64
GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/secpass/secpass-linux-amd64 main.go

# linux arm64
GOOS=linux GOARCH=arm64 go build -o $GOPATH/bin/secpass/secpass-linux-arm64 main.go

# windows amd64
GOOS=windows GOARCH=amd64 go build -o $GOPATH/bin/secpass/secpass-windows-amd64.exe main.go
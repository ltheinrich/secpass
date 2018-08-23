# linux (amd64, arm64)
GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/secpass/secpass-linux-amd64 main/main.go
GOOS=linux GOARCH=arm64 go build -o $GOPATH/bin/secpass/secpass-linux-arm64 main/main.go

# windows amd64
GOOS=windows GOARCH=amd64 go build -o $GOPATH/bin/secpass/secpass-windows-amd64.exe main/main.go
# linux amd64
GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/secpass/secpass-linux-amd64 cmd/secpass/secpass.go

# linux arm64
GOOS=linux GOARCH=arm64 go build -o $GOPATH/bin/secpass/secpass-linux-arm64 cmd/secpass/secpass.go

# linux armv7
GOOS=linux GOARCH=arm GOARM=7 go build -o $GOPATH/bin/secpass/secpass-linux-armv7 cmd/secpass/secpass.go

# windows amd64
GOOS=windows GOARCH=amd64 go build -o $GOPATH/bin/secpass/secpass-windows-amd64.exe cmd/secpass/secpass.go

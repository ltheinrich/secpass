#!/bin/sh
rm -rf target
mkdir target
GOOS=linux GOARCH=amd64 go build -o target/secpass-linux-amd64 cmd/secpass/secpass.go
GOOS=linux GOARCH=arm64 go build -o target/secpass-linux-arm64 cmd/secpass/secpass.go
GOOS=linux GOARCH=arm GOARM=7 go build -o target/secpass-linux-armv7 cmd/secpass/secpass.go
GOOS=windows GOARCH=amd64 go build -o target/secpass-windows-amd64.exe cmd/secpass/secpass.go
tar cfJ target/resources.tar.xz resources

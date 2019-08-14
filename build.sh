#!/bin/sh
rm -rf target
mkdir target
GOOS=linux GOARCH=amd64 go build -o target/secpass-linux-x86_64
GOOS=linux GOARCH=arm64 go build -o target/secpass-linux-aarch64
GOOS=linux GOARCH=arm GOARM=7 go build -o target/secpass-linux-armv7
GOOS=windows GOARCH=amd64 go build -o target/secpass-windows-x86_64.exe
strip target/secpass-linux-x86_64
aarch64-linux-gnu-strip target/secpass-linux-aarch64
arm-linux-gnueabihf-strip target/secpass-linux-armv7
x86_64-w64-mingw32-strip target/secpass-windows-x86_64.exe
tar cfJ target/resources.tar.xz resources

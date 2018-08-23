# linux (amd64, 386, arm64, arm)
GOOS=linux GOARCH=amd64 go build -o ../../../bin/secpass/secpass-linux-amd64 main/main.go
GOOS=linux GOARCH=386 go build -o ../../../bin/secpass/secpass-linux-386 main/main.go
GOOS=linux GOARCH=arm64 go build -o ../../../bin/secpass/secpass-linux-arm64 main/main.go
GOOS=linux GOARCH=arm go build -o ../../../bin/secpass/secpass-linux-arm main/main.go

# windows (amd64, 386)
GOOS=windows GOARCH=amd64 go build -o ../../../bin/secpass/secpass-windows-amd64.exe main/main.go
GOOS=windows GOARCH=386 go build -o ../../../bin/secpass/secpass-windows-386.exe main/main.go

# freebsd (amd64, 386)
GOOS=freebsd GOARCH=amd64 go build -o ../../../bin/secpass/secpass-freebsd-amd64 main/main.go
GOOS=freebsd GOARCH=386 go build -o ../../../bin/secpass/secpass-freebsd-386 main/main.go

# solaris (amd64)
GOOS=solaris GOARCH=amd64 go build -o ../../../bin/secpass/secpass-solaris-amd64 main/main.go

# darwin (amd64, 386)
GOOS=darwin GOARCH=amd64 go build -o ../../../bin/secpass/secpass-darwin-amd64 main/main.go
GOOS=darwin GOARCH=386 go build -o ../../../bin/secpass/secpass-darwin-386 main/main.go

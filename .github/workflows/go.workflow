action "Golang Action" {
  uses = "cedrickring/golang-action@1.3.0"
  args = "go build cmd/secpass/secpass.go"
}

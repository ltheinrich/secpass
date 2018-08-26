package shorts

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"

	// implement postgresql driver
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
)

// LogInfo log only error or everything
var LogInfo = true

// UUID generate random uuid and return it as a string
func UUID() string {
	// generate uuid
	id, _ := uuid.NewV4()
	return id.String()
}

// Check Print error if exists
func Check(err error) {
	// check whether error exists
	if err != nil {
		// print error
		log.Println(err)
	}
}

// ConnectPostgreSQL Connection to postgresql database
func ConnectPostgreSQL(host, port, database, username, password string, ssl string) *sql.DB {
	// open database connection and check for errors
	db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+host+"/"+database+"?sslmode="+ssl)
	Check(err)
	Check(db.Ping())

	// return database connection
	return db
}

// Hash return SHA-3 512 hash string
func Hash(input string) string {
	hasher := sha3.New512()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GenerateKey generate AES key
func GenerateKey(key string) []byte {
	length := len(key)
	switch {
	case length >= 32:
		return []byte(key[:32])
	case length >= 24:
		return []byte(key[:24])
	case length >= 16:
		return []byte(key[:16])
	}

	return GenerateKey(key + "_secpass_key1gen")
}

// Encrypt text with key
func Encrypt(text string, key []byte) string {
	plain := []byte(text)

	block, err := aes.NewCipher(key)
	Check(err)

	cipherText := make([]byte, aes.BlockSize+len(plain))
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	Check(err)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plain)

	return base64.URLEncoding.EncodeToString(cipherText)
}

// Decrypt text with key
func Decrypt(text string, key []byte) string {
	cipherText, err := base64.URLEncoding.DecodeString(text)
	Check(err)

	block, err := aes.NewCipher(key)
	Check(err)

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}

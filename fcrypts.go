package fcrypts

import (
"bytes"
"crypto/aes"
"crypto/cipher"
"crypto/rand"
"crypto/sha1"
"encoding/hex"
"golang.org/x/crypto/pbkdf2"
"io"
"io/ioutil"
"os"
)

func Encrypt(src string, pass []byte) {

	if _, err := os.Stat(src); os.IsNotExist(err) {
		panic(err.Error())
	}

	plaintext, err := ioutil.ReadFile(src)

	if err != nil {
		panic(err.Error())
	}

	key := pass
	nonce := make([]byte, 12)

	// Randomizing the nonce
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherTxt := aesGcm.Seal(nil, nonce, plaintext, nil)

	// Append the nonce to the end of file
	cipherTxt = append(cipherTxt, nonce...)

	f, err := os.Create(src)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(cipherTxt))
	if err != nil {
		panic(err.Error())
	}
}

func Decrypt(source string, password []byte) {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	cipherTxt, err := ioutil.ReadFile(source)

	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := cipherTxt[len(cipherTxt)-12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, cipherTxt[:len(ciphertext)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create(source)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		panic(err.Error())
	}
}


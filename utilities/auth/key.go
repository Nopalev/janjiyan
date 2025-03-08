package auth

import "crypto/rand"

var key []byte

func Init() {
	key = make([]byte, 64)
	rand.Read(key)
}

func getKey() []byte {
	return key
}

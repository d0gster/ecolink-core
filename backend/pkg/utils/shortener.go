package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(url string) string {
	// Combina hash MD5 com timestamp para garantir unicidade
	hasher := md5.New()
	hasher.Write([]byte(url + time.Now().String()))
	hash := hex.EncodeToString(hasher.Sum(nil))
	
	// Pega os primeiros 6 caracteres e converte para base62
	return encodeBase62(hash[:6])
}

func encodeBase62(input string) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
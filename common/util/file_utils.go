package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func CalculateFileHash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		log.Fatal(err)
	}

	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

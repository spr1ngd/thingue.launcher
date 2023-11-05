package util

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

func CalculateFileHash(filePath string) []byte {
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
	return hash
}

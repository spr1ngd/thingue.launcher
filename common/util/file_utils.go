package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"thingue-launcher/common/logger"
)

func CalculateFileHash(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Zap.Error(err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		logger.Zap.Error(err)
	}

	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

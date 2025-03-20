package utils

import (
	"fmt"
	"log"
	"os"
)

func GetCompressedPublicKey() (string, error) {
	log.Println("Checking environment variables...")
	for _, e := range os.Environ() {
		log.Println(e)
	}

	pubKeyPath := "/root/.keys/public_key.txt"
	if _, err := os.Stat(pubKeyPath); os.IsNotExist(err) {
		log.Println("Public key file does not exist:", pubKeyPath)
		return "", fmt.Errorf("public key file not found")
	}

	data, err := os.ReadFile(pubKeyPath)
	if err != nil {
		log.Println("Error reading public key file:", err)
		return "", err
	}

	pubKey := string(data)
	log.Println("Public Key:", pubKey)

	if len(pubKey) < 10 {
		log.Println("Public key is too short:", pubKey)
		return "", fmt.Errorf("invalid public key")
	}

	return pubKey, nil
}

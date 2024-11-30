package auth_coin

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetTxIDFromHex(txHex string) (string, error) {

	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode txHex: %v", err)
	}
	hash1 := sha256.Sum256(txBytes)
	hash2 := sha256.Sum256(hash1[:])

	txID := reverseBytes(hash2[:])

	return hex.EncodeToString(txID), nil
}

func reverseBytes(input []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[len(input)-i-1]
	}
	return output
}

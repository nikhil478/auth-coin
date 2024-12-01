package auth_coin

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/bitcoin-sv/go-sdk/script"
	"github.com/bitcoin-sv/go-sdk/transaction"
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


func PayToAddress(tx *transaction.Transaction, sig []byte, addr string, satoshis uint64) error {
	add, err := script.NewAddressFromString(addr)
	if err != nil {
		return err
	}
	b := make([]byte, 0, 25)
	b = append(b, script.OpDUP, script.OpHASH160, script.OpDATA20)
	b = append(b, add.PublicKeyHash...)
	b = append(b, script.OpEQUALVERIFY, script.OpCHECKSIG)
	s := script.Script(b)
	tx.AddOutput(&transaction.TransactionOutput{
		Satoshis:      satoshis,
		LockingScript: &s,
	})
	return nil
}
package auth_coin

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	ec "github.com/bitcoin-sv/go-sdk/primitives/ec"
	"github.com/bitcoin-sv/go-sdk/script"
	"github.com/bitcoin-sv/go-sdk/transaction"
	"github.com/libsv/go-bk/crypto"
	"github.com/nikhil478/auth-coin/internal/models"
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
	b := make([]byte, 0, 25+len(sig))
	b = append(b, sig...)
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

func SignData(utxo models.UTXO, privKey string) ([]byte, error) {

	priv, err := ec.PrivateKeyFromWif(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key from WIF: %w", err)
	}

	dataToSign :=  crypto.Sha256d([]byte(fmt.Sprintf("%s%d", utxo.TxID, utxo.OutputIndex)))

	signature, err := priv.Sign(dataToSign)
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %w", err)
	}

	sigSerialized := signature.Serialize()
	return sigSerialized, nil
}

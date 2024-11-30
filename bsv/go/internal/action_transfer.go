package auth_coin

import (
	"fmt"

	ec "github.com/bitcoin-sv/go-sdk/primitives/ec"
	"github.com/bitcoin-sv/go-sdk/transaction"
)

func Transfer(txHex string, outputIndex int, holderPrivateKey, issuerPublicKey, issuerPrivateKey, privateKey, destinationAddress string) error {
	
	tx, err := transaction.NewTransactionFromHex(txHex)
	if err != nil {
		return fmt.Errorf("failed to parse transaction: %w", err)
	}

	if outputIndex < 0 || outputIndex >= len(tx.Outputs) {
		return fmt.Errorf("output index %d out of range", outputIndex)
	}

	output := tx.Outputs[outputIndex]

	priv, err := ec.PrivateKeyFromWif(issuerPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to derive private key from WIF: %w", err)
	}

	txID, err := GetTxIDFromHex(txHex)
	if err != nil {
		return fmt.Errorf("failed to get transaction ID from hex: %w", err)
	}

	dataToSign := txID + fmt.Sprintf("%d", outputIndex)

	signature, err := priv.Sign([]byte(dataToSign))
	if err != nil {
		return fmt.Errorf("failed to sign data: %w", err)
	}

	if output.LockingScript.String()[:12] == string(signature.Serialize()) {
		fmt.Printf("Transferring from %s to %s. Amount: %d\n", issuerPublicKey, destinationAddress, output.Satoshis)
	}

	return nil
}


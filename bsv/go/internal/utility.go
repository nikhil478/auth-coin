package auth_coin

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	ec "github.com/bitcoin-sv/go-sdk/primitives/ec"
	"github.com/bitcoin-sv/go-sdk/script"
	"github.com/bitcoin-sv/go-sdk/transaction"
	"github.com/libsv/go-bk/crypto"
	"github.com/nikhil478/auth-coin/internal/models"
)

const (
	LengthOfMultiSigScript = 26
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

func AddOutputWithSignature(tx *transaction.Transaction, addr *string, satoshis uint64, customData *[]byte,) error {
	add, err := script.NewAddressFromString(*addr)
	if err != nil {
		return err
	}
	
	b := make([]byte, 0, 25+len(*customData))
	
	b = append(b, script.OpDUP, script.OpHASH160, script.OpDATA20)
	b = append(b, add.PublicKeyHash...)
	b = append(b, script.OpEQUALVERIFY, script.OpCHECKSIG)

	b = append(b, script.OpNOP)
	b = append(b, *customData...)

	s := script.Script(b)
	tx.AddOutput(&transaction.TransactionOutput{
		Satoshis:      satoshis,
		LockingScript: &s,
	})

	return nil
}

func SignUtxo(utxo *models.UTXO, privKey *string) ([]byte, error) {

	priv, err := ec.PrivateKeyFromWif(*privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key from WIF: %w", err)
	}

	dataToSign :=  crypto.Sha256d([]byte(utxo.TxID))

	signature, err := priv.Sign(dataToSign)
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %w", err)
	}

	sigSerialized := signature.Serialize()
	return sigSerialized, nil
}

func Sha256Hash(data *string) *[]byte {
	hash := crypto.Sha256d([]byte(*data))
	return &hash 
}

func SignData(data *string, privKey *string) (*[]byte, error){
	priv, err := ec.PrivateKeyFromWif(*privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key from WIF: %w", err)
	}
	dataHash := Sha256Hash(data) 
	signature, err := priv.Sign(*dataHash)
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %w", err)
	}

	sigSerialized := signature.Serialize()
	return &sigSerialized, nil
}

func ValidateUtxo(tx *transaction.Transaction,  outputIndex int, issuerPublicKey *string) (bool, error) {
	
	if outputIndex < 0 || outputIndex >= len(tx.Outputs) {
		return false, fmt.Errorf("output index %d out of range", outputIndex)
	}

	output := tx.Outputs[outputIndex]
	messageHex := output.LockingScript.String()[LengthOfMultiSigScript:]
	message, err := hex.DecodeString(messageHex)
	if err != nil {
		return false, err
	}
	sig := message[:23]
	index, err := strconv.Atoi(string(message[23:25]))
    if err != nil {
        return false, err
    }
	inputTxID := tx.Inputs[index].SourceTXID.String()
	hash := Sha256Hash(&inputTxID)
	signature, err := ec.ParseSignature(sig)
	if err != nil {
		return false, err
	}
	publicKey, err := ec.PublicKeyFromString(*issuerPublicKey)
	if err != nil {
		return false, err
	}
	return signature.Verify(*hash, publicKey), nil
}

func ParseUtxo(tx *transaction.Transaction, outputIndex int) (*models.UTXO, error) {

	txID, err := GetTxIDFromHex(tx.Hex())
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction ID from hex: %w", err)
	}

	output := tx.Outputs[outputIndex]

	utxo := models.UTXO{
		TxID: txID,
		OutputIndex: outputIndex,
		Amount: int(output.Satoshis),
		Script: output.LockingScriptHex(),
	}

	return &utxo, nil
}
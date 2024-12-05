package auth_coin

import (
	"fmt"
	"log"

	ec "github.com/bitcoin-sv/go-sdk/primitives/ec"
	"github.com/bitcoin-sv/go-sdk/transaction"
	"github.com/bitcoin-sv/go-sdk/transaction/template/p2pkh"
	"github.com/pkg/errors"
)

func Transfer(txHex string, outputIndex int, holderPrivateKey, issuerPublicKey, issuerPrivateKey, destinationAddress string) (*string, error) {

	tx, err := transaction.NewTransactionFromHex(txHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse transaction: %w", err)
	}

	isValid, err := ValidateUtxo(tx, outputIndex, &issuerPublicKey)
	if !isValid {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("utxo is not valid")
	}

	utxo, err := ParseUtxo(tx, outputIndex)
	if err != nil {
		return nil, err
	}
	signedInfo, err := SignUtxo(utxo , &issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	priv, err := ec.PrivateKeyFromWif(issuerPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key from WIF: %w", err)
	}

	unlockingScriptTemplate, err := p2pkh.Unlock(priv, nil)
	if err != nil {
		return nil, err
	}

	if err := tx.AddInputFrom(
		utxo.TxID,
		uint32(utxo.OutputIndex),
		utxo.Script,
		uint64(utxo.Amount),
		unlockingScriptTemplate,
	); err != nil {
		return nil, err
	}

	err = AddOutputWithSignature(tx, &destinationAddress, uint64(utxo.Amount), &signedInfo)
	if err != nil {
		return nil, err
	}

	if err := tx.Sign(); err != nil {
		log.Fatal(err.Error())
	}

	hex := tx.Hex()

	return &hex, nil
}

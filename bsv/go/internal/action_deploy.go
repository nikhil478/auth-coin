package auth_coin

import (
	"log"

	ec "github.com/bitcoin-sv/go-sdk/primitives/ec"
	"github.com/bitcoin-sv/go-sdk/transaction"
	"github.com/bitcoin-sv/go-sdk/transaction/template/p2pkh"
	"github.com/nikhil478/auth-coin/internal/models"
)

func Deploy(utxo *models.UTXO, issuerPrivateKey *string, holderPrivateKey *string, destinationAddress *string, supply uint64,  additionalData *[]byte, feeUtxo *models.UTXO) (*string , error) {

	tx := transaction.NewTransaction()

	priv, err := ec.PrivateKeyFromWif(*holderPrivateKey)
	if err != nil {
		return nil, err
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

	if err := tx.AddInputFrom(
		feeUtxo.TxID,
		uint32(feeUtxo.OutputIndex),
		feeUtxo.Script,
		uint64(feeUtxo.Amount),
		unlockingScriptTemplate,
	); err != nil {
		return nil, err
	}

	signedInfo, err := SignData(utxo, issuerPrivateKey)
	if err != nil {
		return nil, err
	}

	signedInfo = append(signedInfo, *additionalData...)

	err = AddOutputWithSignature(tx, destinationAddress, supply, &signedInfo)
	if err != nil {
		return nil, err
	}
	if err := tx.Sign(); err != nil {
		log.Fatal(err.Error())
	}

	hex := tx.Hex()

	return &hex, nil
}

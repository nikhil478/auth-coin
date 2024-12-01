package auth_coin

import (
	"log"

	ec "github.com/bitcoin-sv/go-sdk/primitives/ec"
	"github.com/bitcoin-sv/go-sdk/transaction"
	"github.com/bitcoin-sv/go-sdk/transaction/template/p2pkh"
	"github.com/nikhil478/auth-coin/internal/models"
)

func Deploy(utxo models.UTXO, issuerPrivateKey string, holderPrivateKey string, data []string) error {

	tx := transaction.NewTransaction()

	priv, err := ec.PrivateKeyFromWif(holderPrivateKey)

	if err != nil {
		return err
	}

	unlockingScriptTemplate, err := p2pkh.Unlock(priv, nil)
	if err != nil {
		return err
	}

	if err := tx.AddInputFrom(
		utxo.TxID,
		uint32(utxo.OutputIndex),
		utxo.Script,
		uint64(utxo.Amount),
		unlockingScriptTemplate,
	); err != nil {
		log.Fatal(err.Error())
	}

	signedInfo, err := SignData(utxo, issuerPrivateKey)
	if err != nil {
		return err
	}

	err = PayToAddress(tx, signedInfo, "1AdZmoAQUw4XCsCihukoHMvNWXcsd8jDN6", 1000)
	if err != nil {
		return err
	}
	if err := tx.Sign(); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

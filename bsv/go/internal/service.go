package auth_coin

import "github.com/nikhil478/auth-coin/internal/models"

type IVerifier interface {
	Verify(utxo models.UTXO, publicKey string)
}

type ISigner interface {
	Deploy(utxo models.UTXO, privateKey string, data []string)
	Issue(utxo models.UTXO, privateKey string, destinationAddress string)
	Transfer(utxo models.UTXO, privateKey string, destinationAddress string)
	Split(utxo models.UTXO, privateKey string, destinationAddress map[string]int)
	Merge(utxo []models.UTXO, privateKey string, destinationAddress string)
	MergeSplit(utxo []models.UTXO, privateKey string, destinationAddress map[string]int)
	AtomicSwap(utxo1 models.UTXO, utxo2 models.UTXO, destinationAddress1 string, destinationAddress2 string)
}

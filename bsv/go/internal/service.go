package auth_coin

import "github.com/nikhil478/auth-coin/internal/models"

type IVerifier interface {
	Verify(utxo models.UTXO, issuerPublicKey string)
}

type ISigner interface {
	Deploy(utxo models.UTXO, issuerPrivateKey string, data []string)
	Issue(utxo models.UTXO, issuerPublicKey string, issuerPrivateKey string, holderPrivateKey string, destinationAddress string, amount int)
	Transfer(txHex string, outputIndex string, holderPrivateKey string, issuerPublicKey string, issuerPrivateKey string, privateKey string, destinationAddress string)
	Split(txHex string, holderPrivateKey string, issuerPrivateKey string, issuerPublicKey string, destinationAddress map[string]int)
}
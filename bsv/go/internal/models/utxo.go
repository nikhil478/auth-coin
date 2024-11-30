package models 

type UTXO struct {
	TxID string
	OutputIndex int
	Script string
	Amount int
}
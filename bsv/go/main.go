package main

import (
	"fmt"

	auth_coin "github.com/nikhil478/auth-coin/internal"
	"github.com/nikhil478/auth-coin/internal/models"
)

func main() {
	utxo := models.UTXO{
		TxID:        "11b476ad8e0a48fcd40807a111a050af51119099e09283bfa7f3505081a1819d",
		OutputIndex: 0,
		Script:      "76a9144bca0c466925b875875a8e1355698bdcc0b2d45d88ac",
		Amount:      1500,
	}
	signature, err := auth_coin.SignData(utxo, "KznvCNc6Yf4iztSThoMH6oHWzH9EgjfodKxmeuUGPq5DEX5maspS")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Signature: %x\n", signature)
}

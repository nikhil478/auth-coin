package main

import (
	"fmt"

	auth_coin "github.com/nikhil478/auth-coin/internal"
	"github.com/nikhil478/auth-coin/internal/models"
)

func main() {
	utxo := models.UTXO{
		TxID:        "7f308927aa45cf50ddc3b1c31103c7e14d40fa4f00c0e96f726e50a4f61a4a8e",
		OutputIndex: 0,
		Script:      "76a91464229ed55930a6479b9b4731e4e79fbf7e379c8f88ac",
		Amount:      1000,
	}
	hex, err := auth_coin.Deploy(utxo, "L3Fbe9AHwfyypLt2eMGDb6TBunJeh43PvnkJfRdgL1pkF92mZsWd", 
		"L3Fbe9AHwfyypLt2eMGDb6TBunJeh43PvnkJfRdgL1pkF92mZsWd", []string{"Hello World", "Address "})
	if err != nil {
		fmt.Println("Error during deploy:", err.Error())
		return
	}
	auth_coin.Transfer(*hex, 0, "L3Fbe9AHwfyypLt2eMGDb6TBunJeh43PvnkJfRdgL1pkF92mZsWd", "", "L3Fbe9AHwfyypLt2eMGDb6TBunJeh43PvnkJfRdgL1pkF92mZsWd","")
}

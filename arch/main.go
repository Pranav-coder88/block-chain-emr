package main

import (
	"model"
)

func main() {
	//transactions := []model.Transaction{
	//	{Message: "A sent B 5 coins"},
	//	{Message: "A sent C 5 coins"},
	//	{Message: "A sent D 5 coins"},
	//	{Message: "A sent E 5 coins"},
	//	{Message: "A sent F 5 coins"},
	//	{Message: "A sent G 5 coins"}}
	//
	//// Creating a new blockchain with a difficulty of 1.
	//bc := model.NewBlockchain(1)
	//
	//err := bc.AddTransactions(transactions)
	//if err != nil {
	//	fmt.Println("Error adding transaction , ", err)
	//}
	//
	//clock := time.Now()
	//
	//wallet := model.NewWallet()
	//
	//// Processing the pending transactions and rewarding walletID.
	//bc.ProcessPendingTransactions(wallet.GetWalletID())
	//fmt.Println(time.Since(clock))
	//
	//for _, block := range bc.Chain.Links() {
	//	if block.IsValid() {
	//		block.Print()
	//	}
	//
	//	fmt.Println()
	//}

	//model.EncryptBlock()
	model.JSONToEMR()
}

//
//package main
//
//import (
//"fmt"
//"model"
//)
//
//func printChain(bc *model.Blockchain) {
//	for index, block := range bc.Chain.Links() {
//		fmt.Println("\nindex:", index)
//		fmt.Printf("Hashcode: %x\n", block.Hash)
//		fmt.Println("block created on", block.Timestamp)
//		fmt.Println("record created on", block.Transaction.Record.PersonalData.CreatedDate)
//	}
//}
//
//func main() {
//	bc := model.NewBlockchain(2)
//	wallet := model.NewWallet()
//
//	bc.AddTransactions([]model.Transaction{*model.TransactionFromFile("model/dummy.json")})
//
//	bc.ProcessPendingTransactions(wallet.GetWalletID())
//
//	printChain(bc)
//
//}

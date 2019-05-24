package main

import (
	"encoding/json"
	"fmt"
	"log"

	qrcode2 "github.com/skip2/go-qrcode"

	"github.com/algorand/go-algorand-sdk/client/algod"
	"github.com/algorand/go-algorand-sdk/client/kmd"
	"github.com/algorand/go-algorand-sdk/transaction"
)

type NoteFieldType string

const (
	NoteAddStock NoteFieldType = "a"

	NoteSell NoteFieldType = "s"
)

type Note struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`

	Type NoteFieldType `codec:"type"`

	AddStock Stock `codec:"a"`

	SellStock Stock `codec:"s"`
}

func SubmitTxWithNote(blob []byte) {
	kmdClient, err := kmd.MakeClient(kmdAddress, kmdToken)
	if err != nil {
		return
	}

	// Create an algod client
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return
	}

	// Get the suggested transaction params
	txParams, err := algodClient.SuggestedParams()
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}

	// Get suggested fee from algod
	feeResponse, err := algodClient.SuggestedFee()
	if err != nil {
		fmt.Printf("error getting suggested fee: %s\n", err)
		return
	}

	suggestedFee := feeResponse.Fee

	// Make the transaction
	tx, err := transaction.MakePaymentTxn(fisherAddr, fisherAddr, suggestedFee, 1, txParams.LastRound, txParams.LastRound+10, blob, "", txParams.GenesisID)
	if err != nil {
		fmt.Printf("error creating transaction: %s\n", err)
		return
	}

	// Get a wallet handle. The wallet handle is used for things like signing transactions
	// and creating accounts. Wallet handles do expire, but they can be renewed
	initResponse, err := kmdClient.InitWalletHandle(walletID, walletPassword)
	if err != nil {
		fmt.Printf("Error initializing wallet handle: %s\n", err)
		return
	}

	// Extract the wallet handle
	walletHandleToken = initResponse.WalletHandleToken

	// Sign the transaction
	signResponse, err := kmdClient.SignTransaction(walletHandleToken, walletPassword, tx)
	if err != nil {
		fmt.Printf("failed to sign transaction with kmd: %s\n", err)
		return
	}
	fmt.Printf("kmd signed transaction with bytes: %x\n", signResponse.SignedTransaction)

	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(signResponse.SignedTransaction)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction has been broadcasted! ID: %s\n", sendResponse.TxID)
	fmt.Printf("View the transaction here - https://algoexplorer.io/tx/%s\n", sendResponse.TxID)
}

func FormQRCode(from string, amt uint64) []byte {
	// Create an algod client
	algodClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return nil
	}

	// Get the suggested transaction params
	txParams, err := algodClient.SuggestedParams()
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return nil
	}

	tx, err := transaction.MakePaymentTxn(fisherAddr, fisherAddr, 50, amt, txParams.LastRound, txParams.LastRound+100, []byte(""), "", txParams.GenesisID)
	if err != nil {
		log.Printf("Could not make tx for qr code %s", err.Error())
	}

	out, err := json.Marshal(tx)
	qrcode, err := qrcode2.Encode(string(out), qrcode2.Medium, 256)

	return qrcode
}

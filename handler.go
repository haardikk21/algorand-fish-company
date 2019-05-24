package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
)

// AddHarvestHandler handles adding harvest for fisherman
func AddHarvestHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	decoder := json.NewDecoder(req.Body)

	var fishes FishType
	err := decoder.Decode(&fishes)
	log.Printf("%#v", fishes)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	addStock := Stock{
		fisherAddr: {
			fishes,
		},
	}

	addStockBlob := []byte(msgpack.Encode(Note{
		Type:     NoteAddStock,
		AddStock: addStock,
	}))

	SubmitTxWithNote(addStockBlob)

	rw.Write([]byte("Success"))

	go func() {
		time.Sleep(20 * time.Second)
		updateStock()
	}()

	//log.Print(addStockBlob)
	//rw.Write(addStockBlob)
}

// SellHandler handles selling fish to someone else
// it generates and returns a QR code for the multisig transaction
func SellHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	decoder := json.NewDecoder(req.Body)

	var sellStock Stock
	err := decoder.Decode(&sellStock)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	sellStockBlob := []byte(msgpack.Encode(Note{
		Type:      NoteSell,
		SellStock: sellStock,
	}))

	from := ""
	for a := range sellStock {
		from = a
	}

	amt := sellStock[from][0].Amount * sellStock[from][0].Price

	qrcode := FormQRCode(from, uint64(amt))

	qrcodeAsJSON, err := json.Marshal(qrcode)

	SubmitTxWithNote(sellStockBlob)

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(qrcodeAsJSON)

	go func() {
		time.Sleep(20 * time.Second)
		updateStock()
	}()
}

// StockHandler returns the current fisherman stock
func StockHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	currentStockAsJSON, err := json.Marshal(CurrentStock[fisherAddr])
	if err != nil {
		log.Print("ERROR: " + err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(currentStockAsJSON)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/algorand/go-algorand-sdk/client/kmd"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"

	"github.com/algorand/go-algorand-sdk/client/algod"

	"github.com/gorilla/mux"
)

var (
	Port                     = "5000"
	algodAddress             = "http://localhost:8080"
	algodToken               = "3cf8a164de7fdd16dd10948721911cadb5fcea87ba1a0523c0c4147613c2f3cc"
	kmdAddress               = "http://localhost:7833"
	kmdToken                 = "2d517e218cb9773ce564b30c089e5315e4e4281f5b58fc62d8fa04ed2e007299"
	firstRound        uint64 = 1155000
	lastRound         uint64 = 1155500
	fisherAddr               = "fisherman"
	walletPassword           = "abcd"
	walletName               = "fisherman"
	walletID                 = "36336f80ec1ac2b20fe5bf6fb4623bc8"
	walletHandleToken        = ""
	CurrentStock      Stock  = map[string][]FishType{}
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/add", AddHarvestHandler).Methods("POST")
	r.HandleFunc("/sell", SellHandler).Methods("POST")
	r.HandleFunc("/stock", StockHandler).Methods("GET")

	r.Methods("OPTIONS").HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		})
	http.Handle("/", r)

	kmdClient, err := kmd.MakeClient(kmdAddress, kmdToken)
	if err != nil {
		return
	}

	// Create the wallet, if it doesn't already exist
	/*cwResponse, err := kmdClient.CreateWallet(walletName, walletPassword, kmd.DefaultWalletDriver, types.MasterDerivationKey{})
	if err != nil {
		fmt.Printf("error creating wallet: %s\n", err)
		return
	}

	// We need the wallet ID in order to get a wallet handle, so we can add accounts
	walletID = cwResponse.Wallet.ID
	fmt.Printf("Created wallet '%s' with ID: %s\n", cwResponse.Wallet.Name, walletID)*/

	// Get a wallet handle. The wallet handle is used for things like signing transactions
	// and creating accounts. Wallet handles do expire, but they can be renewed
	initResponse, err := kmdClient.InitWalletHandle(walletID, walletPassword)
	if err != nil {
		fmt.Printf("Error initializing wallet handle: %s\n", err)
		return
	}

	// Extract the wallet handle
	walletHandleToken = initResponse.WalletHandleToken

	// Generate a new address from the wallet handle
	genResponse, err := kmdClient.GenerateKey(walletHandleToken)
	if err != nil {
		fmt.Printf("Error generating key: %s\n", err)
		return
	}
	fmt.Printf("Generated address %s\n", genResponse.Address)

	// Extract the wallet address
	fisherAddr = genResponse.Address

	updateStock()

	log.Fatal(http.ListenAndServe(":"+Port, nil))
}

func updateStock() {

	restClient, err := algod.MakeClient(algodAddress, algodToken)
	if err != nil {
		log.Print(os.Stderr, "error making algod client: %v \n", err)
		os.Exit(1)
	}

	//ticker := time.NewTicker(5 * time.Second)
	//for {
	//	select {
	//	case <-ticker.C:
	curRound := firstRound
	finalRound := lastRound
	if curRound > finalRound {
		log.Print("first round %d is after last round %d, exiting\n", curRound, finalRound)
		os.Exit(1)
	}

	CurrentStock = map[string][]FishType{}
	for curRound <= finalRound {
		txns, err := restClient.TransactionsByAddr(fisherAddr, curRound, curRound)
		if err != nil {

		}

		for _, txn := range txns.Transactions {
			if txn.ConfirmedRound != curRound {
				log.Printf("Confirmed round mismatch: found a txn claiming to be confirmed in round %d, in block for round %d", txn.ConfirmedRound, curRound)
				os.Exit(1)
			}

			var note Note
			err = msgpack.Decode(txn.Note, &note)
			if err != nil {
				break
			}

			switch note.Type {
			case NoteAddStock:
				for addr := range note.AddStock {
					if _, ok := CurrentStock[addr]; ok {
						for j := range note.AddStock[addr] {
							found := false
							for i := range CurrentStock[addr] {
								if CurrentStock[addr][i].Species == note.AddStock[addr][j].Species {
									found = true
									CurrentStock[fisherAddr][i].Amount += note.AddStock[addr][j].Amount
								}
							}
							if !found {
								CurrentStock[addr] = append(CurrentStock[addr], note.AddStock[addr]...)
							}
						}
					} else {
						CurrentStock[addr] = note.AddStock[addr]
					}
				}
			case NoteSell:
				for addr := range note.SellStock {
					for j := range note.SellStock[addr] {
						for i := range CurrentStock[fisherAddr] {
							if CurrentStock[fisherAddr][i].Species == note.SellStock[addr][j].Species {
								CurrentStock[fisherAddr][i].Amount -= note.SellStock[addr][j].Amount
							}
						}
					}
				}
			default:
				continue
			}

			//	}
		}
		curRound++
	}
	log.Print("stock collected")
}

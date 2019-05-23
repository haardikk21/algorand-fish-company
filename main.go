package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"

	"github.com/gorilla/mux"
)

var (
	Port               = "5000"
	fisherAddr         = "fisherman"
	CurrentStock Stock = map[string][]FishType{
		fisherAddr: {
			{
				Species: "Tuna",
				Price:   50,
				Image:   "ABC",
				Fishes: []string{
					"1",
					"2",
					"3",
				},
			},
			{
				Species: "Cod",
				Price:   200,
				Image:   "DEF",
				Fishes: []string{
					"4",
					"5",
					"6",
					"7",
					"8",
				},
			},
		},
	}
)

func main() {

	currentStockBlob := []byte(msgpack.Encode(Note{
		Type:     NoteAddStock,
		AddStock: CurrentStock,
	}))

	log.Print(currentStockBlob)

	var note Note
	err := msgpack.Decode(currentStockBlob, &note)
	if err != nil {
		log.Print("ERROR:" + err.Error())
		return
	}

	log.Printf("%#v", note)

	r := mux.NewRouter()

	r.HandleFunc("/add", AddHarvestHandler).Methods("POST")
	r.HandleFunc("/sell", SellHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}

// AddHarvestHandler handles adding harvest for fisherman
func AddHarvestHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var fishes FishType
	err = json.Unmarshal(body, &fishes)
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

	/*found := false
	for i := range CurrentStock[fisherAddr] {
		if CurrentStock[fisherAddr][i].Species == fishes.Species &&
			CurrentStock[fisherAddr][i].Price == fishes.Price {
			found = true
			CurrentStock[fisherAddr][i].Fishes = append(CurrentStock[fisherAddr][i].Fishes, fishes.Fishes...)
		}
	}

	if !found {
		CurrentStock[fisherAddr] = append(CurrentStock[fisherAddr], fishes)
	}

	currentStockJSON, err := json.Marshal(CurrentStock)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	addStockBlob := []byte(msgpack.Encode(Note{
		Type:     NoteAddStock,
		AddStock: addStock,
	}))

	//rw.Header().Set("Content-Type", "application/json")
	//rw.Write(currentStockJSON)

	log.Print(addStockBlob)
	rw.Write(addStockBlob)
}

// SellHandler handles selling fish to someone else
// it generates and returns a QR code for the multisig transaction
func SellHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("ERROR:" + err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var sellStock Stock
	err = json.Unmarshal(body, &sellStock)
	if err != nil {
		log.Print("ERROR: " + err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	sellStockBlob := []byte(msgpack.Encode(Note{
		Type:      NoteSell,
		SellStock: sellStock,
	}))

	log.Print(sellStockBlob)
	rw.Write(sellStockBlob)
}

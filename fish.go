package main

// FishType represents common properties of all fish under the type
type FishType struct {
	Species string `json:"species"`
	Amount  int    `json:"amount"`
	Price   int    `json:"price"`
	Image   string `json:"image"`
}

// Stock is a collection of fishes organized by owner addresses
type Stock map[string][]FishType

package main

// FishType represents common properties of all fish under the type
type FishType struct {
	Species string
	Price   int
	Image   string
	Fishes  []string
}

// Stock is a collection of fishes organized by owner addresses
type Stock map[string][]FishType

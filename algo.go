package main

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

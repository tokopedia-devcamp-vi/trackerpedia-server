package types

import "time"

// Delivery is struct for Delivery tabel
type Delivery struct {
	IDResi  int       `json:"IDResi"`
	Kota    string    `json:"Kota"`
	Status  int       `json:"Status"`
	Ordinal int       `json:"Ordinal"`
	Time    time.Time `json:"Time"`
}

// Order is struct for Order Table
type Order struct {
	IDOrder            int       `json:"IDOrder"`
	NamaItem           string    `json:"NamaItem"`
	Weight             string    `json:"Weight"`
	Status             string    `json:"Status"`
	Time               time.Time `json:"Time"`
	DestinationAddress string    `json:"DestinationAddress"`
	DestinationCity    string    `json:"DestinationCity"`
}

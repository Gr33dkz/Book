package data

import "time"

type Book struct {
	Id          string    `json:"id"`
	Author      string    `json:"author"`
	Quantity    int32     `json:"quantity"`
	Price       float64   `json:"price"`
	ReleaseDate time.Time `json:"releaseDate"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"createdDate"`
}

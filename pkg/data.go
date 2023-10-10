package pkg

import (
	_ "book/docs"
	_ "github.com/swaggo/http-swagger/v2"
	"time"
)

// Book model info
// @Description book Entity
type Book struct {
	Id          string    `json:"id"`
	Author      string    `json:"author"`
	Quantity    int32     `json:"quantity"`
	Price       float64   `json:"price"`
	ReleaseDate time.Time `json:"releaseDate"` // string in UTC format
	Description string    `json:"description"`
	CreatedDate time.Time `json:"createdDate"` // string in UTC format
}

// BookDTO model info
// @Description Entity to create book
type BookDTO struct {
	Author      string    `json:"author"`
	Quantity    int32     `json:"quantity"`
	Price       float64   `json:"price"`
	ReleaseDate time.Time `json:"releaseDate"` // string in UTC format
	Description string    `json:"description"`
}

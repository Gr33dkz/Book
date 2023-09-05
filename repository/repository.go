package repository

import (
	"book/data"
	"database/sql"
	"fmt"
)

//type RepoStore interface {
//	CreateBook(id string, book data.Book) error
//	UpdateBook(book data.Book) error
//	GetBook(id string) *data.Book
//	GetBooks() []data.Book
//	DeleteBook(id string)
//}

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db,
	}
}

func (r *Repo) CreateBook(id string, book data.Book) error {
	fmt.Println("BOOK CREATED")
	return nil
}

func (r *Repo) UpdateBook(book data.Book) error {
	fmt.Println("BOOK UPDATED")
	return nil
}

func (r *Repo) GetBook(id string) *data.Book {
	fmt.Println("GET BOOK")
	return nil
}

func (r *Repo) GetBooks() []data.Book {

	fmt.Println("GET BOOKS")

	return nil
}

func (r *Repo) DeleteBook(id string) {
	fmt.Println("DELETE BOOKS")
}

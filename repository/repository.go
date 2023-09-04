package repository

import (
	"book/data"
	"errors"
)

var Books map[string]data.Book

func InitBook(book data.Book) {
	Books = make(map[string]data.Book)
	Books["111"] = book
}

func CreateBook(id string, book data.Book) error {
	_, exists := Books[id]
	if exists {
		return errors.New("book already exists")
	}
	Books[id] = book
	return nil
}

func UpdateBook(book data.Book) error {
	oldBook, exists := Books[book.Id]
	if !exists {
		return errors.New("book not found")
	}
	oldBook.Price = book.Price
	oldBook.Description = book.Description
	oldBook.ReleaseDate = book.ReleaseDate
	oldBook.Author = book.Author
	oldBook.Quantity = book.Quantity
	Books[book.Id] = oldBook
	return nil
}

func GetBook(id string) *data.Book {
	book, exists := Books[id]
	if !exists {
		return nil
	}
	return &book
}

func GetBooks() []data.Book {

	b := make([]data.Book, 0)
	for _, v := range Books {
		b = append(b, v)
	}

	return b
}

func DeleteBook(id string) {
	delete(Books, id)
}

package service

import (
	"book/data"
	"book/repository"
	"fmt"
)

type Service interface {
	CreateBook(id string, book data.Book) error
	UpdateBook(book data.Book) error
	GetBook(id string) *data.Book
	GetBooks() []data.Book
	DeleteBook(id string)
}

type service struct {
	repo *repository.Repo
}

func New(repo *repository.Repo) Service {
	return &service{
		repo: repo,
	}
}

var Books map[string]data.Book

func (s *service) CreateBook(id string, book data.Book) error {
	fmt.Println("[SERVICE] CreateBook")
	s.repo.CreateBook(id, book)
	//_, exists := Books[id]
	//if exists {
	//	return errors.New("book already exists")
	//}
	//Books[id] = book
	return nil
}

func (s *service) UpdateBook(book data.Book) error {
	fmt.Println("[SERVICE] UpdateBook")
	//oldBook, exists := Books[book.Id]
	//if !exists {
	//	return errors.New("book not found")
	//}
	//oldBook.Price = book.Price
	//oldBook.Description = book.Description
	//oldBook.ReleaseDate = book.ReleaseDate
	//oldBook.Author = book.Author
	//oldBook.Quantity = book.Quantity
	//Books[book.Id] = oldBook
	return nil
}

func (s *service) GetBook(id string) *data.Book {
	fmt.Println("[SERVICE] GetBook")
	//book, exists := Books[id]
	//if !exists {
	//	return nil
	//}
	//return &book
	return nil
}

func (s *service) GetBooks() []data.Book {
	fmt.Println("[SERVICE] GetBooks")
	//b := make([]data.Book, 0)
	//for _, v := range Books {
	//	b = append(b, v)
	//}
	//
	//return b
	return nil
}

func (s *service) DeleteBook(id string) {
	fmt.Println("[SERVICE] GetBooks")
	//delete(Books, id)
}

package service

import (
	"book/data"
	"book/repository"
	"fmt"
)

type Service interface {
	CreateBook(id string, book data.Book) error
	UpdateBook(book data.Book) error
	GetBook(id string) (*data.Book, error)
	GetBooks() []data.Book
	DeleteBook(id string) error
}

type service struct {
	repo *repository.Repo
}

func New(repo *repository.Repo) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateBook(id string, book data.Book) error {
	fmt.Println("[SERVICE] CreateBook")
	err := s.repo.CreateBook(id, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateBook(book data.Book) error {
	fmt.Println("[SERVICE] UpdateBook")
	err := s.repo.UpdateBook(book)
	return err
}

func (s *service) GetBook(id string) (*data.Book, error) {
	fmt.Println("[SERVICE] GetBook")
	book, err := s.repo.GetBook(id)
	return book, err
}

func (s *service) GetBooks() []data.Book {
	books := s.repo.GetBooks()
	return books
}

func (s *service) DeleteBook(id string) error {
	err := s.repo.DeleteBook(id)
	return err
}

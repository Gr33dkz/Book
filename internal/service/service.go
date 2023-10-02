package service

import (
	"book/internal/repository"
	"book/pkg"
	"fmt"
)

type Service interface {
	CreateBook(id string, book pkg.BookDTO) error
	UpdateBook(id string, book pkg.BookDTO) error
	GetBook(id string) (*pkg.Book, error)
	GetBooks() []pkg.Book
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

func (s *service) CreateBook(id string, book pkg.BookDTO) error {
	fmt.Println("[SERVICE] CreateBook")
	err := s.repo.CreateBook(id, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateBook(id string, book pkg.BookDTO) error {
	fmt.Println("[SERVICE] UpdateBook")
	err := s.repo.UpdateBook(id, book)
	return err
}

func (s *service) GetBook(id string) (*pkg.Book, error) {
	fmt.Println("[SERVICE] GetBook")
	book, err := s.repo.GetBook(id)
	return book, err
}

func (s *service) GetBooks() []pkg.Book {
	books := s.repo.GetBooks()
	return books
}

func (s *service) DeleteBook(id string) error {
	err := s.repo.DeleteBook(id)
	return err
}

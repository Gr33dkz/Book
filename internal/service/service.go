package service

import (
	"book/internal/repository"
	"book/pkg"
	log "github.com/sirupsen/logrus"
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
	log.WithFields(log.Fields{
		"methodName": "[CreateBook]",
		"bookId":     id,
	}).Debug("Got request")

	err := s.repo.CreateBook(id, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateBook(id string, book pkg.BookDTO) error {
	log.WithFields(log.Fields{
		"methodName": "[UpdateBook]",
		"bookId":     id,
	}).Debug("Got request")

	err := s.repo.UpdateBook(id, book)
	return err
}

func (s *service) GetBook(id string) (*pkg.Book, error) {
	log.WithFields(log.Fields{
		"methodName": "[GetBook]",
		"bookId":     id,
	}).Debug("Got request")

	book, err := s.repo.GetBook(id)
	return book, err
}

func (s *service) GetBooks() []pkg.Book {
	log.WithFields(log.Fields{
		"methodName": "[GetBooks]",
	}).Debug("Got request")

	books := s.repo.GetBooks()
	return books
}

func (s *service) DeleteBook(id string) error {
	log.WithFields(log.Fields{
		"methodName": "[DeleteBook]",
		"bookId":     id,
	}).Debug("Got request")

	err := s.repo.DeleteBook(id)
	return err
}

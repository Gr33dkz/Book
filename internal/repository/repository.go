package repository

import (
	"book/pkg"
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db,
	}
}

var RecordExists = errors.New("record exists")

func (r *Repo) CreateBook(id string, book pkg.BookDTO) error {
	exists := r.isExists(id)
	log.WithFields(log.Fields{
		"methodName": "CreateBook",
		"isExists":   exists,
	}).Debug()

	if exists {
		return RecordExists
	}
	sqlCreateBook := `INSERT INTO book (id, author, quantity, price, releasedate, description)VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(sqlCreateBook, id, book.Author, book.Quantity, book.Price, book.ReleaseDate, book.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateBook(id string, book pkg.BookDTO) error {
	exists := r.isExists(id)
	log.WithFields(log.Fields{
		"methodName": "CreateBook",
		"isExists":   exists,
	}).Debug()

	if !exists {
		return errors.New("not found")
	}
	stmnt := `update "book" set "author"=$1, "quantity"=$2, "price"=$3, "releasedate"=$4, "description"=$5 where "id"=$6`
	_, err := r.db.Exec(stmnt, book.Author, book.Quantity, book.Price, book.ReleaseDate, book.Description, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetBook(id string) (*pkg.Book, error) {
	stmt := `SELECT id, author, quantity, price, releasedate, description, createddate FROM book WHERE id = $1`
	row := r.db.QueryRow(stmt, id)
	book := &pkg.Book{}

	err := row.Scan(&book.Id, &book.Author, &book.Quantity, &book.Price, &book.ReleaseDate, &book.Description, &book.CreatedDate)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *Repo) GetBooks() []pkg.Book {
	sqlGetBooks := `SELECT * FROM book`
	rows, err := r.db.Query(sqlGetBooks)
	if err != nil {
	}
	defer rows.Close()
	books := make([]pkg.Book, 0)

	for rows.Next() {
		book := pkg.Book{}
		err := rows.Scan(&book.Id, &book.Author, &book.Quantity, &book.Price, &book.ReleaseDate, &book.Description, &book.CreatedDate)
		if err != nil {
			return nil
		}
		books = append(books, book)
	}

	return books
}

func (r *Repo) DeleteBook(id string) error {
	sqlDelete := `DELETE FROM "book" WHERE id = $1`
	_, err := r.db.Exec(sqlDelete, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) isExists(id string) bool {
	query := `SELECT id FROM book WHERE id=$1`
	var idDb string
	row := r.db.QueryRow(query, id)
	err := row.Scan(&idDb)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	}
	return true
}

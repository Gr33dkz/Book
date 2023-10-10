package repository

import (
	"book/pkg"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"testing"
	"time"
)

func Test_CreateBook(t *testing.T) {
	// Test absolute new record
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewRepo(db)

	testId := "generated_test_id"
	mockedTime := time.Now()
	createdSuccess := pkg.BookDTO{
		Author:      "test_author",
		Quantity:    111,
		Price:       11.1,
		ReleaseDate: mockedTime,
		Description: "test_description",
	}

	mock.
		ExpectQuery(`SELECT id FROM book WHERE id=$1`).
		WithArgs(testId).WillReturnError(sql.ErrNoRows)

	mock.ExpectExec(`INSERT INTO book (id, author, quantity, price, releasedate, description)VALUES ($1, $2, $3, $4, $5, $6)`).
		WithArgs(testId, createdSuccess.Author, createdSuccess.Quantity, createdSuccess.Price, createdSuccess.ReleaseDate, createdSuccess.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.CreateBook(testId, createdSuccess)
	if err != nil {
		t.Errorf("[Test_CreateBook] create book assertion failed, got err %v", err)
	}

}

func Test_CreateBook2(t *testing.T) {
	// Test with existed id
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewRepo(db)

	testId := "generated_test_id"
	mockedTime := time.Now()
	createdSuccess := pkg.BookDTO{
		Author:      "test_author",
		Quantity:    111,
		Price:       11.1,
		ReleaseDate: mockedTime,
		Description: "test_description",
	}

	mock.
		ExpectQuery(`SELECT id FROM book WHERE id=$1`).
		WithArgs(testId).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("testId"))

	mock.ExpectExec(`INSERT INTO book (id, author, quantity, price, releasedate, description)VALUES ($1, $2, $3, $4, $5, $6)`).
		WithArgs(testId, createdSuccess.Author, createdSuccess.Quantity, createdSuccess.Price, createdSuccess.ReleaseDate, createdSuccess.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.CreateBook(testId, createdSuccess)
	if err != nil && !errors.Is(RecordExists, err) {
		t.Errorf("[Test_CreateBook2] create book assertion failed, got err %v", err)
	}

}

func Test_GetBookById(t *testing.T) {
	mockedTime := time.Now()
	testId := "test_id"
	expectedBook := pkg.Book{
		Id:          testId,
		Author:      "test_author",
		Quantity:    111,
		Price:       222,
		ReleaseDate: mockedTime,
		Description: "test_description",
		CreatedDate: mockedTime,
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewRepo(db)

	mock.
		ExpectQuery(`SELECT id, author, quantity, price, releasedate, description, createddate FROM book WHERE id = $1`).
		WithArgs(testId).
		WillReturnRows(
			sqlmock.NewRows([]string{`id`, `author`, `quantity`, `price`, `releasedate`, `description`, `createddate`}).
				AddRow(`test_id`, `test_author`, 111, 222, mockedTime, `test_description`, mockedTime))

	uResult, err := repo.GetBook("test_id")
	if err != nil {
		t.Errorf("[Test_GetBookById] test failed %v", err)
	}

	if !reflect.DeepEqual(*uResult, expectedBook) {
		t.Errorf("[Test_GetBookById] assertion book failed got %v, expected%v", uResult, expectedBook)
	}
}

func Test_DeleteBook(t *testing.T) {
	testId := "test_id"

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewRepo(db)

	mock.
		ExpectExec(`DELETE FROM "book" WHERE id = $1`).
		WithArgs(testId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteBook("test_id")
	if err != nil {
		t.Errorf("[Test_DeleteBook] test failed %v", err)
	}

}

func Test_GetBooks(t *testing.T) {
	mockedTime1 := time.Now()
	mockedTime2 := time.Now().Add(time.Minute * 10)
	testId1 := "test_id_1"
	testId2 := "test_id_2"

	expectedBooks := []pkg.Book{
		{
			Id:          testId1,
			Author:      "test_author_1",
			Quantity:    111,
			Price:       111222,
			ReleaseDate: mockedTime1,
			Description: "test_description_1",
			CreatedDate: mockedTime1,
		},
		{
			Id:          testId2,
			Author:      "test_author_2",
			Quantity:    222,
			Price:       222333,
			ReleaseDate: mockedTime2,
			Description: "test_description_2",
			CreatedDate: mockedTime2,
		},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewRepo(db)

	mock.
		ExpectQuery(`SELECT * FROM book`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				`id`, `author`, `quantity`, `price`, `releasedate`, `description`, `createddate`}).
				AddRow(`test_id_1`, `test_author_1`, 111, 111222, mockedTime1, `test_description_1`, mockedTime1).
				AddRow(`test_id_2`, `test_author_2`, 222, 222333, mockedTime2, `test_description_2`, mockedTime2),
		)

	books := repo.GetBooks()

	if !reflect.DeepEqual(expectedBooks, books) {
		t.Errorf("[Test_GetBookById] assertion book failed got %v, expected%v", books, expectedBooks)
	}

}

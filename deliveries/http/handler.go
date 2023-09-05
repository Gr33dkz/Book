package http

import (
	"book/data"
	"book/pkg"
	"book/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Handler struct {
	service service.Service
}

func NewHandler(srv service.Service) *Handler {

	return &Handler{
		service: srv,
	}

}

func (h *Handler) Register() *http.ServeMux {
	mux := http.NewServeMux()
	prefix := "/book"
	mux.Handle(prefix+"/", http.StripPrefix(prefix+"/", http.HandlerFunc(h.handleBooksWithId)))
	mux.Handle(prefix, http.HandlerFunc(h.handleBooks))

	return mux
}

func (h *Handler) handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		rbook := data.Book{}
		bytes, err := readBody(r)
		err = json.Unmarshal(bytes, &rbook)
		if err != nil {
			pkg.UnmarshallError(w)
		}
		err = h.service.CreateBook(rbook.Id, rbook)
		if err != nil {
			pkg.AlreadyExist(w)
		}
		pkg.Accepted(w, "")

	case http.MethodGet:
		var empt []byte
		books := h.service.GetBooks()
		if len(books) == 0 {
			_, err := w.Write(empt)
			if err != nil {
				fmt.Println("error")
			}
		}
		pkg.Response(w, books)

	default:
		pkg.ResponseWithError(w, http.StatusInternalServerError, "Unknown route")
	}
}

func (h *Handler) handleBooksWithId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	switch r.Method {
	case http.MethodDelete:
		h.service.DeleteBook(id)
		pkg.Accepted(w, http.StatusText(http.StatusOK))
	case http.MethodGet:
		b := h.service.GetBook(id)
		if b == nil {
			pkg.ResponseWithError(w, http.StatusNotFound, "Not found")
			return
		}
		pkg.Response(w, b)

	case http.MethodPut:
		rbook := data.Book{}
		bytes, err := readBody(r)
		err = json.Unmarshal(bytes, &rbook)
		if err != nil {
			pkg.UnmarshallError(w)
		}
		err = h.service.UpdateBook(rbook)
		if err != nil {
			pkg.ResponseWithError(w, http.StatusNotFound, "Not found")
			return
		}
		pkg.Accepted(w, http.StatusText(http.StatusOK))
	default:
		pkg.ResponseWithError(w, http.StatusInternalServerError, "Unknown route")
	}
}

func readBody(request *http.Request) ([]byte, error) {
	reqBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return reqBytes, nil
}

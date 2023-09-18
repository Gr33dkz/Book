package v1

import (
	http2 "book/internal/deliveries/http"
	"book/internal/service"
	"book/pkg"
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
		rbook := pkg.Book{}
		bytes, err := readBody(r)
		err = json.Unmarshal(bytes, &rbook)
		if err != nil {
			http2.UnmarshallError(w)
			return
		}
		err = h.service.CreateBook(rbook.Id, rbook)
		if err != nil {
			http2.AlreadyExist(w)
			return
		}
		http2.Accepted(w, "")

	case http.MethodGet:
		var empt []byte
		books := h.service.GetBooks()
		if len(books) == 0 {
			_, err := w.Write(empt)
			if err != nil {
				fmt.Println("error")
			}
		}
		http2.Response(w, books)

	default:
		http2.ResponseWithError(w, http.StatusInternalServerError, "Unknown route")
	}
}

func (h *Handler) handleBooksWithId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	switch r.Method {
	case http.MethodDelete:
		err := h.service.DeleteBook(id)
		if err != nil {
			http2.ResponseWithError(w, http.StatusBadRequest, "DELETE ERROR") // TODO Создать Структуру Ошибки и пробрасывать сообщения
		}
		http2.Accepted(w, http.StatusText(http.StatusOK))
	case http.MethodGet:
		b, err := h.service.GetBook(id)
		if err != nil {
			http2.ResponseWithError(w, http.StatusNotFound, "Not found") // TODO Создать Структуру Ошибки и пробрасывать сообщения
			return
		}
		http2.Response(w, b)

	case http.MethodPut:
		rbook := pkg.Book{}
		bytes, err := readBody(r)
		err = json.Unmarshal(bytes, &rbook)
		if err != nil {
			http2.UnmarshallError(w)
			return
		}
		rbook.Id = id
		err = h.service.UpdateBook(rbook)
		if err != nil {
			http2.ResponseWithError(w, http.StatusNotFound, "Not found")
			return
		}
		http2.Accepted(w, http.StatusText(http.StatusOK))
	default:
		http2.ResponseWithError(w, http.StatusInternalServerError, "Unknown route")
	}
}

func readBody(request *http.Request) ([]byte, error) {
	reqBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return reqBytes, nil
}

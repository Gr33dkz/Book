package v1

import (
	_ "book/docs"
	data "book/internal/deliveries/http"
	"book/internal/service"
	"book/pkg"
	"encoding/json"
	log "github.com/sirupsen/logrus"

	//_ "github.com/go-chi/chi"
	//_ "github.com/swaggo/http-swagger/v2"
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

// @Summary Create Book
// @Tags CRUD
// @Description Endpoint to CREATE book
// @Accept json
// @Produce json
// @Param input body pkg.BookDTO true "using pkg.BookDTO"
// @Response 200 object http.AcceptedResp "Response entity"
// @Failure 400 object http.ErrorMessage "Unmarshall error"
// @Failure 409 object http.ErrorMessage ""Already exists error"
// @Router /book [post]
func (h *Handler) Register() *http.ServeMux {
	mux := http.NewServeMux()
	prefix := "/book"
	mux.Handle(prefix+"/", http.StripPrefix(prefix+"/", http.HandlerFunc(h.handleBooksWithId)))
	mux.Handle(prefix, http.HandlerFunc(h.handleBooks))

	return mux
}

// @Summary Get Books
// @Tags CRUD
// @Description Endpoint to GET-ALL books
// @Accept json
// @Produce json
// @Response 200 object pkg.Book "Response entity"
// @Router /book [get]
func (h *Handler) handleBooks(w http.ResponseWriter, r *http.Request) {
	handlerName := "handleBooks"
	switch r.Method {
	case http.MethodPost:
		rbook := pkg.BookDTO{}
		bytes, err := readBody(r)
		err = json.Unmarshal(bytes, &rbook)
		if err != nil {
			data.UnmarshallError(handlerName, w)
			return
		}
		err = h.service.CreateBook(data.GenerateId(), rbook)
		if err != nil {
			data.AlreadyExist(handlerName, w)
			return
		}
		data.Accepted(w, "")

	case http.MethodGet:
		var empt []byte
		books := h.service.GetBooks()
		if len(books) == 0 {
			_, err := w.Write(empt)
			if err != nil {
				data.MarshallError(handlerName, w)
				return
			}
		}
		data.Response(w, books)

	default:
		log.WithField("handlerName", handlerName)
		data.ResponseWithError(w, http.StatusInternalServerError, "Unknown route")
	}
}

// @Summary Get Book
// @Tags CRUD
// @Description Endpoint to GET book by id
// @Param id path string true "book id"
// @Accept json
// @Produce json
// @Response 200 object pkg.Book "Response entity"
// @Failure 404 object http.ErrorMessage "Not found"
// @Router /book/{id} [get]
func (h *Handler) handleBooksWithId(w http.ResponseWriter, r *http.Request) {
	handlerName := "handleBooksWithId"
	id := r.URL.Path
	switch r.Method {
	case http.MethodDelete:
		err := h.service.DeleteBook(id)
		if err != nil {
			data.ResponseWithError(w, http.StatusBadRequest, "Delete error")
		}
		data.Accepted(w, http.StatusText(http.StatusOK))
	case http.MethodGet:
		b, err := h.service.GetBook(id)
		if err != nil {
			data.ResponseWithError(w, http.StatusNotFound, "Not found")
			return
		}
		data.Response(w, b)
	case http.MethodPut:
		rbook := pkg.BookDTO{}
		bytes, err := readBody(r)
		err = json.Unmarshal(bytes, &rbook)
		if err != nil {
			data.UnmarshallError(handlerName, w)
			return
		}
		err = h.service.UpdateBook(id, rbook)
		if err != nil {
			data.ResponseWithError(w, http.StatusNotFound, "Not found")
			return
		}
		data.Accepted(w, http.StatusText(http.StatusOK))
	default:
		log.WithField("handlerName", handlerName)
		data.ResponseWithError(w, http.StatusInternalServerError, "Unknown route")
	}
}

// @Summary Delete Book
// @Tags CRUD
// @Description Endpoint to DELETE book by id
// @Param id path string true "book id"
// @Accept json
// @Produce json
// @Response 200 object http.AcceptedResp "Response entity"
// @Router /book/{id} [delete]
func readBody(request *http.Request) ([]byte, error) {
	reqBytes, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	return reqBytes, nil
}

// @Summary Update book
// @Tags CRUD
// @Description Endpoint to UPDATE book by id
// @Param id path string true "book id"
// @Accept json
// @Produce json
// @Param input body pkg.BookDTO true "using pkg.BookDTO"
// @Response 200 object http.AcceptedResp "Response entity"
// @Failure 400 object http.ErrorMessage "Unmarshall error"
// @Failure 404 object http.ErrorMessage "Not found"
// @Router /book/{id} [put]
func Mock() {

} // TODO DELETE AFTER MOVING TO ANOTHER FRAMEWORK

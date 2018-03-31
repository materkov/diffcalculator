package diffcalculator

import (
	"github.com/go-chi/chi"
	"net/http"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type calculateRequest struct {
	SourceID string
	Items    map[string]interface{}
}

func (c *calculateRequest) Bind(req *http.Request) error {
	return nil
}

type ApiError struct {
	Error struct {
		Code, Desc string
	}
}

func (e *ApiError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusBadRequest)
	return nil
}

func ErrApi(code, desc string) render.Renderer {
	e := &ApiError{}
	e.Error.Code = code
	e.Error.Desc = desc
	return e
}

var (
	ErrRequestDecodeError = ErrApi("REQUEST_DECODE_ERROR", "Error decoding json request body")
	ErrInternal           = ErrApi("INTERNAL_ERROR", "Internal server error")
)

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	req := calculateRequest{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, ErrRequestDecodeError)
		return
	}

	if err := Calculate(req.SourceID, req.Items); err != nil {
		render.Render(w, r, ErrInternal)
	}

	render.NoContent(w, r)
}

func ServeHTTP() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/calculate", handleCalculate)
	http.ListenAndServe(":8000", r)
}

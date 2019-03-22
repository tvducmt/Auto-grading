package api

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

//Env df
type Env struct {
	DB *sqlx.DB
}

//Index Chữ hoa nếu muôn export , import nếu muốn import
func (e *Env) Index(w http.ResponseWriter, r *http.Request) StatusError {
	return StatusError{0, []byte(`{"result": "OK"}`), nil}
}

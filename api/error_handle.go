package api

import (
	"fmt"
	"net/http"
)

type StatusError struct {
	Code int
	Data []byte
	Err  error
}

func (s StatusError) statusCode() int {
	return s.Code
}

func (s StatusError) dataRes() []byte {
	return s.Data
}

func (s StatusError) errRes() error {
	return s.Err
}

// HandleAPI ...
type HandleAPI func(w http.ResponseWriter, r *http.Request) StatusError

// Reg ...
func (f HandleAPI) Reg(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	if err.dataRes() == nil {
		fmt.Println("statusCode", err.statusCode())
		w.Write([]byte(`{"Error": "Error"}`))
	} else {
		fmt.Println("statusCode", "OKI")
		w.Write([]byte(`{"result": "OK"}`))
	}
}

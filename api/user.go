package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tvducmt/Auto-grading/models"
)

//Register ...
func (e *Env) Register(w http.ResponseWriter, r *http.Request) StatusError {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("into here")
		return StatusError{1, nil, err}
	}
	var acc models.Accout
	err = json.Unmarshal(b, &acc)
	if err != nil {
		fmt.Println("into here1")
		return StatusError{3, nil, err}
	}

	user := models.NewAcount(acc)
	userID, err := user.Add(e.DB)
	if userID == -1 {
		fmt.Println("Eorr", err)
		return StatusError{3, nil, err}
	}
	return StatusError{4, []byte(`{userId: userID}`), nil}
}

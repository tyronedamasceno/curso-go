package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Reading request body failed!"))
		return
	}

	var user user

	if err := json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Error converting user to struct"))
	}

	fmt.Println(user)

}

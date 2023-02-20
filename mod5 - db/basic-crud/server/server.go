package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var user user

	if err := json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Error converting user to struct"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to db"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		w.Write([]byte("Error creating statement"))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		w.Write([]byte("Error executing statement"))
		return
	}

	insertedId, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Error retrieving inserted id"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User successfully created! ID: %d", insertedId)))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to db"))
		return
	}
	defer db.Close()

	lines, err := db.Query("SELECT * FROM users")
	if err != nil {
		w.Write([]byte("Error querying the users"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user

		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error while scanning user"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error encoding users to json"))
		return
	}
}

func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error converting param to int"))
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to db"))
		return
	}
	defer db.Close()

	line, err := db.Query("SELECT * FROM users WHERE id = ?", ID)

	if err != nil {
		w.Write([]byte("Error retrieving user"))
		return
	}

	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning user"))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error converting user to json"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error converting param to int"))
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var user user

	if err := json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Error converting user to struct"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to db"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		w.Write([]byte("Error creating statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Error updating user"))
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error converting param to int"))
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to db"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		w.Write([]byte("Error preparing statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Error deleting user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

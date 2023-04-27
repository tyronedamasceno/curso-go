package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userOnDB, err := repository.RetrieveUserByEmail(user.Email)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPwd(userOnDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userOnDB.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
}

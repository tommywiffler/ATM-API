package handlers

import (
	"atm-api/database"
	"atm-api/models"
	"atm-api/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to read request body, see error: %v", err))
		return
	}
	err = json.Unmarshal(body, &login)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to unmarshal body see error: %v", err))
		return
	}
	var user models.User
	loginResult, err := user.UserLogin(database.DB, login.ID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to login, see error: %v", err))
		return
	}

	responses.JSON(w, http.StatusOK, loginResult)
}

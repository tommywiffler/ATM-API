package handlers

import (
	"atm-api/database"
	"atm-api/models"
	"atm-api/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["id"]
	var account models.Account
	balanceResult, err := account.GetBalance(database.DB, owner)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to get balance, see error: %v", err))
		return
	}

	responses.JSON(w, http.StatusOK, balanceResult)

}

func Deposit(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to read request body, see error: %v", err))
		return
	}
	err = json.Unmarshal(body, &account)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to unmarshal body see error: %v", err))
		return
	}

	balanceResult, err := account.Deposit(database.DB, account.Number)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to process deposit, see error: %v", err))
		return
	}

	responses.JSON(w, http.StatusOK, balanceResult)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to read request body, see error: %v", err))
		return
	}
	err = json.Unmarshal(body, &account)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to unmarshal body see error: %v", err))
		return
	}

	balanceResult, err := account.Withdraw(database.DB, account.Number)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to process withdrawal, see error: %v", err))
		return
	}

	responses.JSON(w, http.StatusOK, balanceResult)
}

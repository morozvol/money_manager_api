package controllers

import (
	"github.com/gorilla/mux"
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/store"
	"github.com/morozvol/money_manager_api/tools"
	"net/http"
	"strconv"
)

func GetAccounts(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userId := params["userId"]
		id, err := strconv.Atoi(userId)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}
		accounts, err := store.Account().FindByUserId(id)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}
		tools.SendResponse(rw, http.StatusOK, "success", accounts)
	}
}

func GetAccount(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		sId := params["accountId"]
		id, err := strconv.Atoi(sId)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}
		account, err := store.Account().Find(id)
		if err != nil {
			tools.SendError(rw, http.StatusInternalServerError, err)
			return
		}
		tools.SendResponse(rw, http.StatusOK, "success", account)
	}
}

func CreateAccount(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var account model.Account
		if !tools.Decode(rw, r, &account) {
			return
		}
		if !tools.Validate(rw, &account) {
			return
		}
		storeErr := store.Account().Create(&account)
		if storeErr != nil {
			tools.SendError(rw, http.StatusInternalServerError, storeErr)
			return
		}
		tools.SendResponse(rw, http.StatusCreated, "success", account)
	}
}

func DeleteAccount(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		sId := params["accountId"]
		id, err := strconv.Atoi(sId)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}
		account, err := store.Account().Find(id)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}
		account.IsActual = false
		editAccount(store, account)
	}
}

func EditAccount(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var newAccount model.Account
		if !tools.Decode(rw, r, &newAccount) {
			return
		}
		if !tools.Validate(rw, &newAccount) {
			return
		}
		editAccount(store, &newAccount)
	}
}

func editAccount(store store.Store, newAccount *model.Account) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		account, err := store.Account().Find(newAccount.Id)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}

		account.IsActual = newAccount.IsActual
		account.Balance = newAccount.Balance
		account.Name = newAccount.Name

		err = store.Account().Update(account)
		if err != nil {
			tools.SendError(rw, http.StatusInternalServerError, err)
			return
		}
		tools.SendResponse(rw, http.StatusOK, "success", nil)
	}
}

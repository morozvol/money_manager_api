package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/store"
	"github.com/morozvol/money_manager_api/tools"
	"net/http"
	"strconv"
)

func CreateOperation(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var operation model.Operation

		if !tools.Decode(rw, r, &operation) {
			return
		}
		if !tools.Validate(rw, &operation) {
			return
		}
		storeErr := store.Operation().Create(&operation)
		if storeErr != nil {
			tools.SendError(rw, http.StatusInternalServerError, storeErr)
			return
		}
		tools.SendResponse(rw, http.StatusCreated, "success", operation)
	}
}

func GetOperation(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		operationId := params["operationId"]
		var operation *model.Operation

		id, err := strconv.Atoi(operationId)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}
		operation, err = store.Operation().Find(id)
		if err != nil {
			tools.SendError(rw, http.StatusNotFound, err)
			return
		}
		tools.SendResponse(rw, http.StatusOK, "success", *operation)
	}
}

func EditOperation(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		operationId := params["operationId"]
		id, err := strconv.Atoi(operationId)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, errors.New(fmt.Sprintf("error: не удалость превести параметр operationId(%s) к int", operationId)))
			return
		}
		var operation model.Operation
		if !tools.Decode(rw, r, &operation) {
			return
		}
		operation.Id = id
		if !tools.Validate(rw, &operation) {
			return
		}
		storeErr := store.Operation().Update(&operation)
		if storeErr != nil {
			tools.SendError(rw, http.StatusInternalServerError, err)
			return
		}
		tools.SendResponse(rw, http.StatusOK, "success", operation)
	}
}

func DeleteOperation(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		operationId := params["operationId"]
		id, err := strconv.Atoi(operationId)
		if err != nil {
			tools.SendError(rw, http.StatusBadRequest, err)
			return
		}
		err = store.Operation().Delete(id)
		if err != nil {
			tools.SendError(rw, http.StatusInternalServerError, err)
			return
		}
		tools.SendResponse(rw, http.StatusOK, "success", nil)
	}
}

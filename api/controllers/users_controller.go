package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/ncjain/RestAPI/api/models"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)
	err = server.DB.Debug().Create(&user).Error
	if err != nil {
		data := struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	errr := json.NewEncoder(w).Encode(&user)
	if err != nil {
		fmt.Fprintf(w, "%s", errr.Error())
	}
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users := []models.User{}
	err := server.DB.Debug().Model(&user).Limit(100).Find(&users).Error
	if err != nil {
		data := struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	errr := json.NewEncoder(w).Encode(&users)
	if err != nil {
		fmt.Fprintf(w, "%s", errr.Error())
	}
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

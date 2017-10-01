package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/valentyn88/app"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService app.UserService
}

func NewUserHandler(userService app.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userService.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user app.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = h.userService.Create(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	decoder.Decode(&data)

	if err = h.userService.Update(id, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

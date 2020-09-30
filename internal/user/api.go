package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	srv Service
}

// RegisterHandlers add routes of blogs
func RegisterHandlers(routers *mux.Router, service Service) {
	h := handler{srv: service}
	routers.HandleFunc("/user/save", h.saveUser).Methods("POST")
	routers.HandleFunc("/user/{user}", h.getUser).Methods("GET")
}

func (h handler) getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user"]
	user, err := h.srv.Get(context.TODO(), userID)
	if err != nil {
		fmt.Println("error fething user", err)
	}
	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in json convert", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (h handler) saveUser(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	name := r.FormValue("name")
	email := r.FormValue("email")
	data := CreateUserRequest{
		UserID: userID,
		Name:   name,
		Email:  email,
	}
	user, err := h.srv.Create(context.TODO(), data)
	if err != nil {
		fmt.Println("error creating user", err)
	}

	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in json convert", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

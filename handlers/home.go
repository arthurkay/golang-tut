package handlers

import (
	"db_proj/db"
	"db_proj/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserDetails struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func New(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)

	userDetails := UserDetails{}
	if err != nil {
		fmt.Fprintf(w, "We ran into an error: %v", err)
	}

	err = json.Unmarshal(req, &userDetails)
	if err != nil {
		fmt.Fprintf(w, "We ran into an error: %v", err)
	}
	userData := &models.User{
		Name:     userDetails.Name,
		Password: userDetails.Password,
	}

	dbInstance, err := db.DB()

	if err != nil {
		fmt.Fprintf(w, "We ran into an error: %v", err)
	}

	db.CreateUser(dbInstance, userData)

}

func First(w http.ResponseWriter, r *http.Request) {
	dbInstance, err := db.DB()
	if err != nil {
		fmt.Fprintf(w, "We ran into an error: %v", err)
	}
	userDetails := db.GetFirstUser(dbInstance)
	fmt.Fprintf(w, userDetails.Name)
}

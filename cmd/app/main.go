package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/valentyn88/app"
	httpApp "github.com/valentyn88/app/http"
	"log"
	"net/http"
)

var cache *app.Cache

func init() {
	cache = app.NewCache()
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	basicUser := app.NewBasicUser(cache)
	userHandler := httpApp.NewUserHandler(basicUser)

	r := mux.NewRouter()
	r.HandleFunc("/user/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/user/new", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", userHandler.UpdateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

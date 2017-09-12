package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type App struct {
	Router  *mux.Router
	Session *mgo.Session
}

func (a *App) Init(user, password, dbname string) {

	s, err := mgo.Dial("127.0.0.1:7778/test")

	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.SetMode(mgo.Monotonic, true)

	a.Session = s.Copy()
	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) Run(addr string) {
	fmt.Println("listen")
	log.Fatal(http.ListenAndServe(":7777", a.Router))
}

func (a *App) initializeRoutes() {
	// User
	a.Router.HandleFunc("/users", a.postUsers).Methods("POST")
	a.Router.HandleFunc("/users/{id}", a.getUsers).Methods("GET")
	// Company
	a.Router.HandleFunc("/companies", a.postCompanies).Methods("POST")
	a.Router.HandleFunc("/companies/{id}", a.getCompanies).Methods("GET")
}

func (a *App) postUsers(w http.ResponseWriter, r *http.Request) {
	var u User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request - new user - "+err.Error())
		return
	}

	if err := u.createUser(a.Session); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)

}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	var u User

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u, err := u.getUser(id, a.Session)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)

}

func (a *App) postCompanies(w http.ResponseWriter, r *http.Request) {
	var c Company
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&c); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request - new company - "+err.Error())
		return
	}

	if err := c.createCompany(a.Session); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, c)

}

func (a *App) getCompanies(w http.ResponseWriter, r *http.Request) {
	var c Company

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	c, err := c.getCompany(id, a.Session)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, c)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

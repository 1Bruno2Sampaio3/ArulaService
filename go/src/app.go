package main

import (
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
	//only a test
	//defer s.Close()

	s.SetMode(mgo.Monotonic, true)

	a.Session = s
	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) Run(addr string) {
	defer a.Session.Close()
	log.Fatal(http.ListenAndServe(":7777", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {

}

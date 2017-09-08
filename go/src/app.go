package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type App struct {
	Router  *mux.Router
	Session *mgo.Session
}

func (a *App) Init(user, password, dbname string) {

	s, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.SetMode(mgo.Monotonic, true)

	a.Session = s
	a.Router = mux.NewRouter()

}

func (a *App) Run(addr string) {

}

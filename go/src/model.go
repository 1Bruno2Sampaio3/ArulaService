package main

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
)

// Main structs

type User struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	CPF         string       `json:"cpf"`
	Address     string       `json:"address"`
	Birth       time.Time    `json:"date"`
	Skills      []Skill      `json:"skills"`
	Experiences []Experience `json:"experiences"`
	Jobs        []Job        `json:"jobs"`
}

type Company struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	CPF      string  `json:"cpf"`
	Address  string  `json:"address"`
	Area     []Skill `json:"skills"`
}

//Substructs

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Experience struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}

type Job struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Desc     string    `json:"desc"`
	Date     time.Time `json:"date"`
	Finished bool      `json:"finished"`
}

func (u *User) createUser(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *User) getUser(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *User) updateUser(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *User) deleteUser(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func getUsers(db *mgo.Session, start, count int) error {
	return errors.New("Not implemented")
}

func (u *Company) createCompany(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *User) getCompany(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *User) updateCompany(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *User) deleteCompany(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func getCompanies(db *mgo.Session, start, count int) error {
	return errors.New("Not implemented")
}

func (u *Job) getJob(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *Job) updateJob(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func (u *Job) deleteJob(db *mgo.Session) error {
	return errors.New("Not implemented")
}

func getJobs(db *mgo.Session, start, count int) error {
	return errors.New("Not implemented")
}

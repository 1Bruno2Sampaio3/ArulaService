package main

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Main structs

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	CPF         string        `json:"cpf"`
	Address     string        `json:"address"`
	Birth       time.Time     `json:"birth"`
	Skills      []Skill       `json:"skills"`
	Experiences []Experience  `json:"experiences"`
	Jobs        []Job         `json:"jobs"`
}

type Company struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	CPF      string        `json:"cpf"`
	Address  string        `json:"address"`
	Area     []Skill       `json:"skills"`
}

//Substructs

type Skill struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name"`
}

type Experience struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name"`
	Occupation string        `json:"occupation`
	Start      time.Time     `json:"start"`
	End        time.Time     `json:"end"`
}

type Job struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title    string        `json:"title"`
	Desc     string        `json:"desc"`
	Date     time.Time     `json:"date"`
	Finished bool          `json:"finished"`
}

func (u *User) createUser(db *mgo.Session) error {

	u.ID = newId()

	c := db.DB("arula-test").C("users")
	if err := c.Insert(u); err != nil {
		return err
	}

	return nil

}

func (u *User) getUser(id string, db *mgo.Session) (User, error) {

	result := User{}

	if err := validId(id); err != nil {
		return result, err
	}

	c := db.DB("arula-test").C("users")
	if err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result); err != nil {
		return result, err
	}

	return result, nil

}

func (u *User) updateUser(id string, db *mgo.Session) error {

	if err := validId(id); err != nil {
		return err
	}

	c := db.DB("arula-test").C("users")
	if err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, u); err != nil {
		return err
	}

	u.ID = bson.ObjectIdHex(id)

	return nil

}

func (u *User) deleteUser(id string, db *mgo.Session) error {

	if err := validId(id); err != nil {
		return err
	}

	c := db.DB("arula-test").C("users")
	if err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		return err
	}

	return nil

}

func (c *Company) createCompany(db *mgo.Session) error {

	c.ID = newId()

	cc := db.DB("arula-test").C("companies")
	if err := cc.Insert(c); err != nil {
		return err
	}

	return nil

}

func (c *Company) getCompany(id string, db *mgo.Session) (Company, error) {

	result := Company{}

	if err := validId(id); err != nil {
		return result, err
	}

	cc := db.DB("arula-test").C("companies")
	if err := cc.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result); err != nil {
		return result, err
	}

	return result, nil

}

func getUsers(db *mgo.Session, start, count int) error {
	return errors.New("Not implemented")
}

func (u *User) getCompanies(db *mgo.Session) error {
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

func newId() bson.ObjectId {
	return bson.NewObjectId()
}

func validId(id string) error {

	if !bson.IsObjectIdHex(id) {
		return errors.New("Id is not valid")
	}

	return nil

}

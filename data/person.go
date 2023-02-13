package data

import "log"

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Zipcode string `json:"zipcode"`
}

type People []*Person

type PersonDB struct {
	db PersonDAO
	l  log.Logger
}

func NewPersonDB(db PersonDAO, l log.Logger) *PersonDB {
	return &PersonDB{db, l}
}

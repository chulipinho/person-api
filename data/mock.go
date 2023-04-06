package data

import (
	"errors"
)

type Mock struct{}

func NewMock() *Mock {
	return &Mock{}
}

func (db *Mock) Get() ([]Person, error) {
	return people, nil
}

func (db *Mock) GetById(id int) (*Person, error) {
	p, _ := findPersonById(id)
	if p == nil {
		return nil, errors.New("Could not find person with given id")
	}

	return p, nil
}

func (db *Mock) Post(p Person) error {
	if p.Id == 0 {
		p.Id = getLastId() + 1
	}
	people = append(people, p)

	return nil
}

func (db *Mock) Put(id int, p Person) error {
	person, i := findPersonById(id)
	if person == nil {
		db.Post(p)
		return nil
	}

	p.Id = id
	people[i] = p
	return nil
}

func (db *Mock) Delete(id int) error {
	_, i := findPersonById(id)
	if i == -1 {
		return errors.New("Element not found")
	}

	people[i] = people[len(people)-1]
	people = people[:len(people)-1]
	return nil
}

func findPersonById(id int) (*Person, int) {
	for i, p := range people {
		if p.Id == id {
			return &p, i
		}
	}

	return nil, -1
}

func getLastId() int {
	len := len(people)
	return people[len-1].Id
}

var people = []Person{
	{
		Id:      1,
		Name:    "Adler",
		Email:   "adler@example.com",
		Zipcode: "30240360",
	},
	{
		Id:      2,
		Name:    "Pedro",
		Email:   "pedro@example.com",
		Zipcode: "35701210",
	},
}

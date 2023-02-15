package data

type PersonDAO interface {
	Get() ([]Person, error)
	GetById(int) (*Person, error)
	Post(Person) error
	Put(int, Person) error
	Delete(int) error
}

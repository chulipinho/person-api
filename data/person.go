package data

type Person struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Zipcode string `json:"zipcode"`
}

type People []*Person

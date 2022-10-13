package entity

//Person object for REST(CRUD)

type Address struct {
	Address string `json:"address"`
}

type Contact struct {
	phoneNumber int    `json:"phone"`
	email       string `json:"email"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Address
	Contact
}

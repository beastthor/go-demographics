package entity

//Person object for REST(CRUD)

type Address struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}

type Contact struct {
	Id          int `json:"id"`
	PhoneNumber int `json:"phone"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string
	Contact   string
}

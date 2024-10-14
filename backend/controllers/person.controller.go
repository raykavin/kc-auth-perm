package controllers

import (
	"encoding/json"
	"net/http"
)

type (
	PersonController struct{}
	Person           struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
		Document  string `json:"document"`
		Address   string `json:"address"`
		Phone     string
	}
)

var people = []Person{
	{
		FirstName: "John",
		LastName:  "Doe",
		Age:       32,
		Document:  "123-45-6789",
		Address:   "123 Elm St, Springfield",
		Phone:     "(555) 123-4567",
	},
	{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       28,
		Document:  "987-65-4321",
		Address:   "456 Oak Ave, Metropolis",
		Phone:     "(555) 987-6543",
	},
	{
		FirstName: "Michael",
		LastName:  "Johnson",
		Age:       40,
		Document:  "321-54-9876",
		Address:   "789 Pine Rd, Gotham",
		Phone:     "(555) 321-9876",
	},
	{
		FirstName: "Emily",
		LastName:  "Davis",
		Age:       25,
		Document:  "654-32-1987",
		Address:   "101 Maple Blvd, Star City",
		Phone:     "(555) 654-3219",
	},
}

func (*PersonController) List(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(people)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(res)
}

func NewPersonController() *PersonController {
	return &PersonController{}
}

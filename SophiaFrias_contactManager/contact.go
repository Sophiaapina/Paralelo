package main

// Contact is the main model for the contact manager.
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

const DefaultDataFile = "contacts.json"

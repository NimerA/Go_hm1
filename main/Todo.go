package main

import "time"

//Todo is a model
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos is a ordered collection of Todo
type Todos []Todo

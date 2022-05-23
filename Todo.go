package main

import "time"

type Todo struct {
	Id        int       `json:id`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	Sales     Nested    `json:"sales"`
}

type Nested struct {
	FirstName string `json:"firstName"`
	LirstName string `json:"lastName"`
	Age       int    `json:"age"`
}
type Todos []Todo

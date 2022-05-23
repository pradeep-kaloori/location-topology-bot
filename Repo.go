package main

import "fmt"

var currentId int

var todos Todos

// Give us some seed data
func init() {
	//RepoCreateTodo(Todo{Name: "Write presentation"})
	//RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoCreateTodo(t []Todo) []Todo {
	var newTodo []Todo
	fmt.Println("This is the curretn value of crrentId", currentId)
	// t.Id = currentId + 1
	// todos = append(todos, t)
	// return t
	for _, todo := range t {
		todo.Id = currentId + 1
		currentId++
		fmt.Println("This is the curretn value of crrentId in for loop", currentId)
		newTodo = append(newTodo, todo)
	}
	return newTodo
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

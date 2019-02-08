package main

import (
	"log"
	"net/http"

	"github.com/go_example_vanilla/homepage"
	todo "github.com/go_example_vanilla/todo_list"
)

func main() {
	http.HandleFunc("/", homepage.Homepage)
	http.HandleFunc("/todo_list", todo.EnterPage)
	http.HandleFunc("/todo_list/add", todo.AddItem)
	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)

}

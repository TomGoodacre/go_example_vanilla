package main

import (
	"log"
	"net/http"

	"github.com/go_example_vanilla/homepage"
	todo "github.com/go_example_vanilla/todo_list"
)

func main() {
	http.HandleFunc("/", homepage.Homepage)
	http.HandleFunc("/todo_list", todo.ViewList)
	http.HandleFunc("/todo_list/add_item", todo.EnterAddItem)
	http.HandleFunc("/todo_list/add_new_item", todo.AddItem)
	http.HandleFunc("/todo_list/delete", todo.RemoveItem)

	//Serve static files with StripPrefix, in case the directory changes.
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)

}

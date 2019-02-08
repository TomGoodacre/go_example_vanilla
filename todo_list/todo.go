package todo

import (
	"net/http"
)

type todoItem struct {
	DateTime string
	Notes    string
}

//Variables for the main todo page
type Variables struct {
	ToDoList []todoItem
}

//EnterPage creates the main todo page, on navigation to the page
func EnterPage(w http.ResponseWriter, r *http.Request) {

}

//AddItem adds a todo item to the list, and displays it
func AddItem(w http.ResponseWriter, r *http.Request) {

}

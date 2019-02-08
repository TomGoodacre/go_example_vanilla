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

//CreatePage creates the main todo page
func CreatePage(w http.ResponseWriter, r *http.Request) {

}

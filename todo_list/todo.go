package todo

import (
	"log"
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
	err := r.ParseForm()
	if err != nil {
		log.Println("error in parsing form: ", err)
	}

	//If everything is working as it should, r.Form is
	//map[date: [value], time: [value], notes: [value]]
	//(? In what format are the values from date/time input in this form ?)
}

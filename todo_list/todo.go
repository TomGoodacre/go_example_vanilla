package todo

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type todoItem struct {
	DateTime string
	Notes    string
}

//Variables for the main todo page
type Variables struct {
	TodoList []todoItem
}

var items []todoItem

//EnterPage creates the main todo page, on navigation to the page
func EnterPage(w http.ResponseWriter, r *http.Request) {
	//Store the todo list in the PageVars struct sent to todo.html
	PageVars := Variables{
		TodoList: items,
	}

	//Parse the todo_list html
	t, err := template.ParseFiles("todo_list/todo.html")
	//Check for errors
	if err != nil {
		log.Println("template parsing error: ", err)
	}

	//Execute the todo_list page
	err = t.Execute(w, PageVars)
	//Check for errors
	if err != nil {
		log.Println("template execute error: ", err)
	}
}

//AddItem adds a todo item to the list, and displays it
func AddItem(w http.ResponseWriter, r *http.Request) {
	//Populate the form with values passed in the request.
	err := r.ParseForm()
	//Check for parsing error.
	if err != nil {
		log.Println("error in parsing form: ", err)
	}

	//r.Form is now map[date: [value], time: [value], notes: [value]]

	//The date and time received from the form.
	formDate := r.Form.Get("date") + " " + r.Form.Get("time")
	//The layout of the time/date in formDate.
	layout := "2006-01-02 15:04"
	//Parse into a Time, in accordance with the above layout.
	date, err := time.Parse(layout, formDate)
	//Check for parsing error.
	if err != nil {
		log.Println("date parsing error: ", err)
	}

	//Create a new todo, with a formatted date/time.
	newTodo := todoItem{
		DateTime: date.Format("02-Jan (Mon) 15:04"),
		Notes:    r.Form.Get("notes"),
	}

	//Insert the new todo into the start of the item list.
	items = append([]todoItem{newTodo}, items...)

	//Store the todo list in the PageVars struct sent to todo.html
	PageVars := Variables{
		TodoList: items,
	}

	//Parse the todo_list html
	t, err := template.ParseFiles("todo_list/todo.html")
	//Check for parsing error
	if err != nil {
		log.Println("template parsing error: ", err)
	}

	//Execute the todo_list page
	err = t.Execute(w, PageVars)
	//Check for execution error
	if err != nil {
		log.Println("template execute error: ", err)
	}
}

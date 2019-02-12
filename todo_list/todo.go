package todo

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type todoItem struct {
	dateAdded        time.Time
	dateTodo         time.Time
	DisplayDateAdded string
	DisplayTodoDate  string
	Notes            string
}

//Variables for the main todo page
type Variables struct {
	TodoList []todoItem
}

var items []todoItem

//ViewList creates the main todo page, on navigation to the page
func ViewList(w http.ResponseWriter, r *http.Request) {
	//Sort the todo list by todo time.
	_sortByTodoTime(items)

	//Store the todo list in the PageVars struct sent to todo.html
	PageVars := Variables{
		TodoList: items,
	}

	//Parse the todo_list html
	t, err := template.ParseFiles("todo_list/todo_list.html")
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

	notes := r.Form.Get("notes")

	if err == nil && notes != "" {
		//Create a new todo, with a formatted date/time.
		newTodo := todoItem{
			dateAdded:        time.Now(),
			dateTodo:         date,
			DisplayDateAdded: time.Now().Format("02-Jan (Mon) 15:04"),
			DisplayTodoDate:  date.Format("02-Jan (Mon) 15:04"),
			Notes:            notes,
		}

		//Insert the new todo into the start of the item list.
		items = append([]todoItem{newTodo}, items...)

		log.Println("todo item added")
	}

	//Store the todo list in the PageVars struct sent to todo.html
	PageVars := Variables{
		TodoList: items,
	}

	//Parse the todo_list html
	t, err := template.ParseFiles("todo_list/todo_add.html")
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

func RemoveItem(w http.ResponseWriter, r *http.Request) {

}

func _sortByTimeAdded(unsorted []todoItem) {
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i].dateAdded.Before(unsorted[j].dateAdded)
	})
}

func _sortByTodoTime(unsorted []todoItem) {
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i].dateTodo.Before(unsorted[j].dateTodo)
	})
}

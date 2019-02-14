package todo

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type todoItem struct {
	ID               int
	dateAdded        time.Time
	dateTodo         time.Time
	DisplayDateAdded string
	DisplayTodoDate  string
	Notes            string
}

//struct にfuncを追加出来ます。classのメソッドみたい
func (item todoItem) Urgency() string {
	timeUntil := time.Until(item.dateTodo)
	switch {
	case timeUntil < time.Hour*24:
		return "urgent"
	default:
		return "normal"
	}

}

//Variables for the main todo page
type Variables struct {
	TodoList []todoItem
}

//For the remove json
type removeData struct {
	deleteIndex int
}

var idCounter int
var items []todoItem

//ViewList creates the main todo page, on navigation to the page
func ViewList(w http.ResponseWriter, r *http.Request) {
	//Sort the todo list by todo time.
	_sortByTodoTime(items)

	//Serve todo_list.html
	_showTodoListHTML(&w)
}

//EnterAddItem is used for navigating to the add_item page.
func EnterAddItem(w http.ResponseWriter, r *http.Request) {
	_showAddItemHTML(&w)
}

//AddItem adds a todo item to the list, and displays it
func AddItem(w http.ResponseWriter, r *http.Request) {
	log.Println("adding item...")

	//Parse the values from the Request into a todoItem struct,
	//and add it to the todoItem list
	_addItemToList(r)

	//Serve the html from the page html template, with the new item list.
	_showAddItemHTML(&w)

}

//RemoveItem removes a todo item from the list of items, and reserves the html
func RemoveItem(w http.ResponseWriter, r *http.Request) {
	ajaxData := r.FormValue("delete_id")
	log.Println("ajax data: ", ajaxData)

	removeID, err := strconv.Atoi(ajaxData)
	if err != nil {
		log.Println("parsing error from remove index data")
		return
	}

	removeIndex := -1
	for i, item := range items {
		if item.ID == removeID {
			removeIndex = i
			break
		}
	}

	if removeIndex >= 0 {
		//Delete an element, the golang way...
		items = append(items[:removeIndex], items[removeIndex+1:]...)
	}

	log.Println("item deleted")

	//Reserve the todo_list.html
	_showTodoListHTML(&w)

}

func _showTodoListHTML(w *http.ResponseWriter) {
	log.Println("showing todo_list.html")

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
	err = t.Execute(*w, PageVars)
	//Check for errors
	if err != nil {
		log.Println("template execute error: ", err)
	}
}

func _addItemToList(r *http.Request) {
	//Populate the form with values passed in the request.
	err := r.ParseForm()
	//Check for parsing error.
	if err != nil {
		log.Println("error in parsing form: ", err)
	}

	//r.Form is now map[
	//					date: [string (format: 2006-01-02)],
	//					time: [string (format: 15:04)],
	//					notes: [string],
	//				   ]

	formDate := r.Form.Get("date") + " " + r.Form.Get("time")
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
			ID:               idCounter,
			dateAdded:        time.Now(),
			dateTodo:         date,
			DisplayDateAdded: time.Now().Format("02-Jan (Mon) 15:04"),
			DisplayTodoDate:  date.Format("02-Jan (Mon) 15:04"),
			Notes:            notes,
		}

		idCounter++

		//Insert the new todo into the start of the item list.
		items = append([]todoItem{newTodo}, items...)

		log.Println("todo item added")
	}
}

func _showAddItemHTML(w *http.ResponseWriter) {
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
	err = t.Execute(*w, PageVars)
	//Check for execution error
	if err != nil {
		log.Println("template execute error: ", err)
	}
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

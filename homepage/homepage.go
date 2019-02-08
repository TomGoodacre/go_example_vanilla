package homepage

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

//HelloWorld The entry point for the homepage
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	bytesWritten, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println("bytes written: ", bytesWritten, "error: ", err)
	}
}

//Variables are the date and time for the homepage
type Variables struct {
	Date string
	Time string
}

//Homepage uses the homepage.html template and HomepageVariables to create the home page
func Homepage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	//Store the date and time in a struct which will be used in the homepage.html template
	PageVars := Variables{
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04:05"),
	}

	//Parse the homepage.html
	t, err := template.ParseFiles("homepage/homepage.html")
	//Check for errors
	if err != nil {
		log.Println("template parsing error: ", err)
	}
	//Execute the template (i.e. write the template, using the provided writer and the
	//context data provided in the variables struct passed in this function)
	err = t.Execute(w, PageVars)
	if err != nil {
		log.Println("template executing error: ", err)
	}

}

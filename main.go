package main

import (
	"fmt"
	"net/http"

	"github.com/go_example_vanilla/homepage"
)

func main() {
	http.HandleFunc("/", homepage.Homepage)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error = ", err)
}

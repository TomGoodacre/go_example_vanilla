package main

import (
	"fmt"
	"net/http"

	"github.com/go_raw_example/homepage"
)

func main() {
	http.HandleFunc("/", homepage.Homepage)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error = ", err)
}

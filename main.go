package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	bytesWritten, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println("bytes written: ", bytesWritten, "error: ", err)
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error = ", err)
}

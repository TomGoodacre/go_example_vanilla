package homepage

import (
	"fmt"
	"net/http"
)

//HelloWorld The entry point for the homepage
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	bytesWritten, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println("bytes written: ", bytesWritten, "error: ", err)
	}
}

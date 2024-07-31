package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	if request.Method != http.MethodGet {
		http.Error(writer, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(writer, "hello!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "Parse form error: %v", err)
		return
	}

	fmt.Fprintf(writer, "POST request successful")

	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "Name: %v \n", name)
	fmt.Fprintf(writer, "Address: %v \n", address)
}

// SIMPLE WEB SERVER IN GO
// Santiago Garcia Arango

package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHelloHandler(w http.ResponseWriter, r *http.Request) {
	// If someone uses this path by mistake, return error 404
	if r.URL.Path != "/hello" {
		message := "404 not found"
		http.Error(w, message, http.StatusNotFound)
		fmt.Printf(message)
		return
	}
	if r.Method != "GET" {
		message := "Method not allowed"
		http.Error(w, message, http.StatusNotFound)
		fmt.Printf(message)
		return
	}
	fmt.Fprintf(w, "HELLO SANTI'S FRIENDS!\n")
	fmt.Printf("GET method at [%v] path\n", "hello")
}

func myFormHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET") && (r.Method != "POST") {
		message := "Method not allowed"
		http.Error(w, message, http.StatusNotFound)
		fmt.Printf(message)
		return
	}
	if r.Method == "GET" {
		// For "GET", it should be with the static "/form.html" path, not "/form"
		http.Redirect(w, r, "/form.html", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		message := fmt.Sprintf("ParseForm() error: %v\n", err)
		fmt.Fprintf(w, message)
		fmt.Printf(message)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")
	message := fmt.Sprintf("--> Name: %s\n--> Age: %s\n", name, age)
	fmt.Fprintf(w, message)
	fmt.Printf("POST request successfully handled:\n%s\n", message)
}

func main() {
	// These lines will enable "static" folder for "index.html" and "form.html"
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Serve "/hello" and "/form" paths
	http.HandleFunc("/hello", myHelloHandler)
	http.HandleFunc("/form", myFormHandler)

	fmt.Printf("Starting golang backend server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

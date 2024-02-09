package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/login", FormHandler)
	fmt.Printf("Listening on port 8080\n")

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/login.html")
		return
	}

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Fprintf(w, "Email = %s\n", email)
		fmt.Fprintf(w, "Password = %s\n", password)
		return
	}
	
}

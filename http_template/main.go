package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", WelcomeHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))


	http.ListenAndServe(":8085", nil)
}

type User struct {
	Name        string
	nationality string //unexported field.
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("http_template/welcomeform.html")
		check(err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		myUser := User{}
		myUser.Name = r.Form.Get("entered_name")
		myUser.nationality = r.Form.Get("entered_nationality")
		t, err := template.ParseFiles("http_template/welcomeresponse.html")
		check(err)
		t.Execute(w, myUser)
	}
}

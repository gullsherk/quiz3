package main

import (
    "html/template"
    "net/http"
)


func mainPage(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    data := "I AM THE BEST "
    tmpl.Execute(w, data)
}

func pakistan(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/pakistan.html"))
    data := "I AM THE BEST "
    tmpl.Execute(w, data)
}

func southkorea(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/southkorea.html"))
    data := "I AM THE BEST "
    tmpl.Execute(w, data)
}

func france(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/france.html"))
    data := "I AM THE BEST "
    tmpl.Execute(w, data)
}

func message(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/message.html"))
    data := "I AM THE BEST "
    tmpl.Execute(w, data)
}

func main() {
    http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

    http.HandleFunc("/", mainPage)
    http.HandleFunc("/pakistan", pakistan)
    http.HandleFunc("/southkorea", southkorea)
    http.HandleFunc("/france", france)
    http.HandleFunc("/message", message)
    
    http.ListenAndServe(":8080", nil)
}
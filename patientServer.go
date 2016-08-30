package main 

import (
    "fmt"
    "net/http"
    "html/template"
)

type Patient struct {
    Name string
   // Email string
   // Phone string
    //Passwd string
}


func adminHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/admin/"):] 
    fmt.Fprintf(w, "Hello %s", title)
}

func userHandler(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("dashboard.html")
    t.Execute(w, nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request){
    name := r.FormValue("Name")
    p := &Patient{Name: name}

    fmt.Fprintf(w, "Hello %s", p.Name)


    }



func main() {
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/dashboard/", userHandler)
    http.HandleFunc("/register/", registerHandler)
    http.ListenAndServe(":8080", nil)
}



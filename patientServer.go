package main 

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Patient struct {
    Name string
   Email string
    Phone string
    Passwd string
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
    name, email, phone, passwd := r.FormValue("Name"), r.FormValue("Email"), r.FormValue("Phone"), r.FormValue("Password")
    p := &Patient{Name: name, Email: email, Phone: phone, Passwd: passwd}
    fmt.Fprintf(w, "%s%s%s%s", p.Name, p.Email, p.Phone, p.Passwd)


    }



func main() {
    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/customer")
    if err != nil {
            log.Fatal(err)
    }
    defer db.Close()
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/dashboard/", userHandler)
    http.HandleFunc("/register/", registerHandler)
    http.ListenAndServe(":8080", nil)
}



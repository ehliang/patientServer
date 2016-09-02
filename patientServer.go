package main 

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    "github.com/gorilla/securecookie"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Patient struct {
    Name string
   Email string
    Phone string
    Passwd string
}

var cookieHandler = secureCookie.New(
    securecookie.GenerateRandomKey(64), 
    securecookie.GenerateRandomKey(32))

var router = mux.NewRouter()


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

	db.Exec

    }

func userLoginHandler(w http.ResponseWriter, r *http.Request){
    setSession()
}

func userLogoutHandler(w http.ResponseWriter, r *http.Request){
    clearSession()
}

func setSession()
{
}

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/customer")
    if err != nil {
            log.Fatal(err)
    }
    defer db.Close()
    http.HandleFunc("/admin/", adminHandler).Methods("POST")
    http.HandleFunc("/login/", userLoginHandler).Methods("POST")
    http.HandleFunc("/logout/", userLogoutHandler).Methods("POST")
    http.HandleFunc("/dashboard/", userHandler).Methods("POST")
    http.HandleFunc("/register/", registerHandler).Methods("POST")
    http.ListenAndServe(":8080", nil)
}



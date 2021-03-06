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

var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64), 
    securecookie.GenerateRandomKey(32))

var router = mux.NewRouter()


func adminHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/admin/"):] 
    fmt.Fprintf(w, "Hello %s", title)
}

func registerHandler(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("register.html")
    t.Execute(w, nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request){

    userName := isLoggedIn(r)
    if userName != ""{
        t, _ := template.ParseFiles("dashboard.html")
        t.Execute(w, nil)
    }else{
        fmt.Fprint(w, "Access Denied")    
    }
}

func confirmationHandler(w http.ResponseWriter, r *http.Request){
    name, email, phone, passwd := r.FormValue("Name"), r.FormValue("Email"), r.FormValue("Phone"), r.FormValue("Password")
    p := &Patient{Name: name, Email: email, Phone: phone, Passwd: passwd}
    fmt.Fprintf(w, "%s%s%s%s", p.Name, p.Email, p.Phone, p.Passwd)


    }

func indexHandler(w http.ResponseWriter, r *http.Request){

    t, _ := template.ParseFiles("index.html")
    t.Execute(w, nil) 
}



func userLoginHandler(w http.ResponseWriter, r *http.Request){
    name := r.FormValue("Name")
    pass := r.FormValue("Password") 
    redirectTarget:="/abcd/"
    if name=="abc" && pass=="abc"{
         setSession(name, w) 
         redirectTarget="/dashboard"        
    } 
    http.Redirect(w, r, redirectTarget, 302)    
}

func userLogoutHandler(w http.ResponseWriter, r *http.Request){
    clearSession(w)
    http.Redirect(w, r,"/index/", 302)
}

func setSession(userName string, w http.ResponseWriter){
    value := map[string]string{
         "name":userName, 
    }
    if encoded, err := cookieHandler.Encode("session", value); err == nil{
         cookie := &http.Cookie{
             Name: "session", 
             Value: encoded, 
             Path: "/", 
        }
        http.SetCookie(w, cookie) 
    }
}

func isLoggedIn(r *http.Request) (userName string) {
     if cookie, err := r.Cookie("session"); err == nil{
         cookieValue := make(map[string]string) 
             if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
                 userName = cookieValue["name"]
             }
     }
     return userName
}



func clearSession(w http.ResponseWriter){
    cookie := &http.Cookie{
        Name: "session", 
        Value: "", 
        Path: "/", 
        MaxAge:-1,
    }
    http.SetCookie(w, cookie) 
}

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/customer")
    if err != nil {
            log.Fatal(err)
    }
    defer db.Close()
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/index/", indexHandler)
    http.HandleFunc("/login/", userLoginHandler)
    http.HandleFunc("/logout/", userLogoutHandler)
    http.HandleFunc("/dashboard/", dashboardHandler)
    http.HandleFunc("/register/", registerHandler)
    http.HandleFunc("/confirm/", confirmationHandler)
    http.ListenAndServe(":8080", nil)
}



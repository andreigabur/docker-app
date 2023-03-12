package main

import (
    "fmt"
    "net/http"
    "docker-app/database"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "Hi There!\nGo server with auto-reload here!\n\n")
    fmt.Fprintf(w, "The list of users from Database\n")

    database.ConnectDB()
    rows, _ := database.DB.Raw("select name, email from users").Rows()
    defer rows.Close()
    for rows.Next() {
        var name string 
        var email string
        rows.Scan(&name, &email)
        fmt.Fprintf(w, "name: %v, email: %v\n", name, email)
    }
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8080", nil)
}
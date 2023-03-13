package main

import (
    "fmt"
    "net/http"
    "encoding/json"

    "docker-app/database"
)

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getUsers(w http.ResponseWriter, req *http.Request) {

    enableCors(&w)

    type User struct {
        Name string
        Email string
    }

    type JsonResponse struct {
        Type    string `json:"type"`
        Data    []User `json:"data"`
        Message string `json:"message"`
    }

    var users []User

    database.ConnectDB()
    rows, _ := database.DB.Raw("select name, email from users").Rows()
    defer rows.Close()
    for rows.Next() {
        var name string 
        var email string
        rows.Scan(&name, &email)
        users = append(users, User{Name: name, Email: email})
    }

    var response = JsonResponse{Type: "success", Data: users}
    json.NewEncoder(w).Encode(response)
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hi There!\nGo server with auto-reload here!\n\n")
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

    http.HandleFunc("/getusers", getUsers)

    http.ListenAndServe(":8080", nil)
}
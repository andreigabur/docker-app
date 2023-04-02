package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "encoding/json"
    "github.com/Shopify/sarama"

    "docker-app/database"
)

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Go server started on port 8080! Hi")
}

func kafka(w http.ResponseWriter, req *http.Request) {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true

    //verbose debugging (comment this line to disabled verbose sarama logging)
    sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

    brokers := []string{"kafka:9092"} // replace with the address of your Kafka broker
    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := producer.Close(); err != nil {
            panic(err)
        }
    }()

    message := &sarama.ProducerMessage{
        Topic: "my-topic",
        Value: sarama.StringEncoder("hello world!"),
    }

    partition, offset, err := producer.SendMessage(message)
    if err != nil {
        panic(err)
    }

    fmt.Fprintf(w, "Message sent to partition %d at offset %d\n", partition, offset)
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

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/kafka", kafka)

    http.HandleFunc("/getusers", getUsers)

    http.ListenAndServe(":8080", nil)
}
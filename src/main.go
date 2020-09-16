package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/lib/pq"
)

func main() {

	dbSettings := database.Settings{
		User: "postgres",
		Pass: "P!r0genk0123",
		Host: "localhost",
		Port: "5432",
		Name: "phone_list",
	}

	err := database.Connect(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	// 127.0.0.1/test
	http.HandleFunc("/test", test)

	http.HandleFunc("/get", get)
	http.HandleFunc("/get_all", getAll)
	http.HandleFunc("/update", update)
	http.HandleFunc("/add", add)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request url" + r.RequestURI)
	fmt.Println("Method " + r.Method)

	fmt.Println("Params:")

	for k, v := range r.URL.Query() {
		fmt.Println(k + " = " + v[0])
	}

}

// GET /get?id
func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		// http.Error(w, "405 Method is not alloved", 405)
		http.Error(w, "405 Method is not alloved", http.StatusMethodNotAllowed)
		return
	}
}

// GET /getAll
func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 Method is not alloved", http.StatusMethodNotAllowed)
		return
	}
}

// Update /update?id=1&phone=12345678
func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "UPDATE" {

		http.Error(w, "405 Method is not alloved", http.StatusMethodNotAllowed)
		return
	}
}

// PUT /add?id=1&phone=12345678
func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "405 Method is not alloved", http.StatusMethodNotAllowed)
		return
	}
}

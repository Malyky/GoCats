package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	name, age := returnValuesTest(5)
	fmt.Println(name, age)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/cats", cats)
	log.Println("Starting server :8080")
	http.ListenAndServe(":8080", mux)
}

func hello(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Hello, world!")
	w.Write([]byte("Hello, Louisa!"))
}

func cats(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Hello, world!")
	w.Write([]byte("Beautiful Cat pictures here"))
}

func returnValuesTest(number int) (name string, age int) {

	name = "abc"
	age = age + 2
	return
}

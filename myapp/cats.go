package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"myapp/config"
	"myapp/myapp/api/resource/cat"
	"net/http"
)

func main() {
	name, age := returnValuesTest(7)
	fmt.Println("TEST")
	fmt.Println(name, age)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/cats", cats)
	log.Println("Starting server :8080")
	fs := http.FileServer(http.Dir("./ressources"))
	mux.Handle("/", fs)
	c := config.New()

	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", "localhost", c.DB.Username, c.DB.Password, c.DB.DBName, c.DB.Port) //TODO
	fmt.Println(dbString)
	db, err := gorm.Open(postgres.Open(dbString))

	if err != nil {
		log.Fatal("DB connection start failure")
		return
	}

	catRepo := cat.NewCatRepository(db)
	cats, _ := catRepo.List()

	for _, cat := range cats {
		fmt.Printf("Cat Name: %s, Age: %d\n", cat.Name, cat.Age)
	}
	fmt.Print(cats)

	http.ListenAndServe(":8080", mux)

}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, Louisa!"))
}

func cats(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Hello, world!")
	w.Write([]byte("Beautiful Cat pictures here"))
}

func returnValuesTest(number int) (name string, age int) {

	name = "qwe"
	age = age + 2
	return
}

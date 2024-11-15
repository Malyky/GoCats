package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"myapp/config"
	"myapp/myapp/api/resource/cat"
	"myapp/myapp/api/router"
	"net/http"
)

var db gorm.DB

func main() {
	db := getDb()
	name, age := returnValuesTest(7)
	fmt.Println(name, age)

	// Define http Endpoints
	mux := http.NewServeMux()
	mux.HandleFunc("/cat/{name}", hello)
	mux.HandleFunc("/findmycat/{name}", findmycat)
	fs := http.FileServer(http.Dir("./ressources"))
	mux.Handle("/", fs)

	// Cat Repository
	catRepo := cat.NewCatRepository(db)
	cats, _ := catRepo.List()

	for _, cat := range cats {
		fmt.Printf("Cat Name: %s, Age: %d\n", cat.Name, cat.Age)
	}
	fmt.Print(cats)

	log.Println("Starting server :8080")
	http.ListenAndServe(":8080", mux)

	// Other rest router
	r := router.New(db)
	fmt.Print(r)

}

func hello(w http.ResponseWriter, req *http.Request) {
	path := req.PathValue("name")

	fmt.Print(path)
	w.Write([]byte("Hello FedExDay - " + path + ", to the Cats app!"))
}

func findmycat(w http.ResponseWriter, req *http.Request) {
	path := req.PathValue("name")

	// Cat Repository
	catRepo := cat.NewCatRepository(getDb())
	cat, _ := catRepo.Find(path)
	fmt.Print(cat.Name)

	fmt.Print(path)
	w.Write([]byte(fmt.Sprintf("Cat has age %d and name %s", cat.Age, cat.Name)))
}

func cats(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Hello, world!")
	w.Write([]byte("Beautiful Cat pictures here"))
}

func returnValuesTest(number int) (name string, age int) {
	name = "hello"
	age = age + 2
	fmt.Print("Hallo :)")
	return
}

func getDb() (db *gorm.DB) {
	c := config.New()
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", "localhost", c.DB.Username, c.DB.Password, c.DB.DBName, c.DB.Port) //TODO
	fmt.Println(dbString)
	fmt.Println(dbString)

	// Database Connection
	db, err := gorm.Open(postgres.Open(dbString))
	if err != nil {
		log.Fatal("DB connection start failure")
		return
	}
	return
}

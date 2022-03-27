package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type App struct {
	Router mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {

	a.DB, _ = gorm.Open("sqlite3", "test.db")
	a.DB.AutoMigrate(&Todo{})

	a.Router = *mux.NewRouter()
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/todos", a.Todos).Methods("GET")
	a.Router.HandleFunc("/todos", a.CreateTodo).Methods("POST")
	a.Router.HandleFunc("/todos/{id}", a.Todo).Methods("PUT")

}
func (a *App) Run(port string) {

	fmt.Println("Listening at " + port)

	err := http.ListenAndServe(port, &a.Router)
	if err != nil {
		log.Fatal(err)
	}
}
func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
func (a *App) Todos(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "İsmail"},
		Todo{Name: "Furkan"},
	}
	respondWithJson(w, http.StatusOK, todos)
	// json.NewEncoder(w).Encode(todos)
}

func (a *App) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	decoder := json.NewDecoder(r.Body) //get body
	if err := decoder.Decode(&todo); err != nil {
		respondWithError(w, http.StatusOK, "Failed")
	}
	a.DB.Create(&todo)
	respondWithJson(w, http.StatusOK, todo)
}
func (a *App) Todo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Single todo page ", vars["id"])
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

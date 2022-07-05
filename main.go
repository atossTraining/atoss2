package main

import (
	"ex4/entity"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var person entity.Person

func main() {

	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		router.Get("/person", GetPerson)
		router.Post("/person", CreatePerson)
	})
	router.Route("/v2", func(r chi.Router) {
		router.Get("/person", GetPerson)
		router.Post("/person", CreatePerson2)
	})
	http.ListenAndServe(":8080", router)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r.Method)
	value := fmt.Sprintf("Person got is %v", person)
	fmt.Fprintln(w, value)
	fmt.Fprintln(w, "GET method")

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	//body := r.Body
	//decodedBody, _ := io.ReadAll(body)
	//json.Unmarshal(decodedBody, &person)
	person = entity.Person{
		Name:   "name",
		Gender: "gender",
		ID:     "id",
		Age:    129,
	}
	value := fmt.Sprintf("Person just created is %v", person)
	fmt.Fprintln(w, value)
	fmt.Fprintln(w, "POST method 1")

}

func CreatePerson2(w http.ResponseWriter, r *http.Request) {

	//body := r.Body
	//decodedBody, _ := io.ReadAll(body)
	//json.Unmarshal(decodedBody, &person)
	person = entity.Person{
		Name:   "name",
		Gender: "gender",
		ID:     "id",
		Age:    229,
	}
	value := fmt.Sprintf("Person just created is %v", person)
	fmt.Fprintln(w, value)
	fmt.Fprintln(w, "POST method 2")

}

//IT ALWAYS STARTS LIKE THIS
package main

//THE STANDARD GOLANG SHENANIGANS TO IMPORT PACKAGES
import (
	"encoding/json"
	"log" //TO USE THE log.Fatal FUNCTION
	"math/rand"
	"net/http" //TO HANDLE HTTP REQUEST
	"strconv"

	"github.com/gorilla/mux" //TO CREATE ROUTES. PEOPLE TEND TO USE THIS PACKAGE
)

//Book MODEL. BEHAVES LIKE A CLASS
type Book struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`

	//THE ASTERISC IS CAUSE THE AUTHOR IS GOING TO HAVE HIS OWN MODEL (STRUCT). WE REFERENCE OTHER STRUCTS USING THIS SYNTAX
	Author *Author `json:"author"`
}

//Author MODEL. BEHAVES LIKE A CLASS
type Author struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

//INITIATE BOOK VAR AS A SLICE (ARRAY WITH MUTABLE SIZE)
var books []Book

//getBook FUNCTION
//EVERY FUNCTION THAT IS A ROUT HANDLER NEED TO HAVES THESE PROPERTIES. SO, BASICALY JUST COPY PASTE IT WHEN NEEDED.
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books) //w IS FOR WRITING IN THE SCREEN, r IS FOR RECIEVING A RESPONSE FROM THE FRONTEND (PARAMS)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //GET PARAMS
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{}) //OUTPUTS ALL THE BOOKS
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //MOCK ID
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//LIKE IN C, IT ALWAYS NEED A MAIN FUNCTION TO START THINGS OFF
func main() {

	//MUX IS USED TO CREATE THE BACKEND ROUTES
	r := mux.NewRouter()

	//MOCK DATA(SOME VARIABLES FOR TESTING PURPOSES). THE TRUE DATA WILL BE SENT FROM THE FRONTEND
	books = append(books,
		Book{ID: "1",
			Isbn:  "4487",
			Title: "Book One",
			Author: &Author{
				Firstname: "Xand",
				Lastname:  "ao"}})

	books = append(books,
		Book{ID: "2",
			Isbn:  "44234",
			Title: "Book Two",
			Author: &Author{
				Firstname: "Vi",
				Lastname:  "tu"}})

	//CREATING THE ROUTES
	//THE FIRST PARAMETER IS THE NAME OF THE ROUT "/something/whatever", THE SECOND, IS THE FUNCTION IT CALLS.
	//THEN WE USE .Methods(GET OR POST OR DELETE OR WHATEVER) TO SPECIFY THE METHOD WE ARE USING IN THIS ROUT
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Println("Server Online!")
	//SETS THE SERVER UP AT PORT 3035. YOU CAN USE WHATEVER PORT YOU WANT
	//THE log.Fatal FUNCTION, IS USED JUST TO LOG THE ERRORS IF THEY HAPPEN
	log.Fatal(http.ListenAndServe(":3035", r))

}

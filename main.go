package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Movie struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// to ensure there is data to start with, append a Movie struct/object
	movies = append(movies, Movie{
		ID: "1",
		ISBN: "1401223176",
		Title: "The Batman",
		// & reference of the director object address 
		Director: &Director{
			Firstname: "Matt",
			Lastname: "Reeves",
		},
	})
	movies = append(movies, Movie{
		ID: "2",
		ISBN: "2401223177",
		Title: "Top Gun",
		// & reference of the director object address 
		Director: &Director{
			Firstname: "Tony",
			Lastname: "Scott",
		},
	})
	// req to /movies, respond with getMovies function 
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// START SERVER 
	fmt.Printf("Starting server at PORT 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
	
}
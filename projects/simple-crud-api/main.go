// SIMPLE CRUD API
// Santiago Garcia Arango

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Director *Director `json:"director"`
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
}

type Director struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func (m Movie) showMovieDetails() {
	// TODO
	return
}

type Message struct {
	Message string `json:"message"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// Create fake data (this will be database in the future)
	loadSampleMovies()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Initializing server at port: 9000\n")
	log.Fatal(http.ListenAndServe(":9000", r))

}

func loadSampleMovies() {

	movies = append(
		movies,
		Movie{
			Director: &Director{Name: "Gore", LastName: "Verbinski"},
			Id:       "1",
			Title:    "Pirates of the Caribbean: The Curse of the Black Pearl",
			Year:     2003,
		},
		Movie{
			Director: &Director{Name: "Andrew", LastName: "Niccol"},
			Id:       "2",
			Title:    "In Time",
			Year:     2011,
		},
		Movie{
			Director: &Director{Name: "Morten", LastName: "Tyldum"},
			Id:       "3",
			Title:    "The Imitation Game",
			Year:     2014,
		},
		Movie{
			Director: &Director{Name: "Marc", LastName: "Webb"},
			Id:       "4",
			Title:    "Gifted",
			Year:     2017,
		},
		Movie{
			Director: &Director{Name: "Christopher", LastName: "Nolan"},
			Id:       "5",
			Title:    "Interstellar",
			Year:     2014,
		},
	)

	fmt.Printf("Loaded movies are: %v\n", movies)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_movie Movie
	_ = json.NewDecoder(r.Body).Decode(&new_movie)
	// This for is to guarantee that new movies ids are not repeated
	for {
		repeatedId := false
		new_movie.Id = strconv.Itoa(rand.Intn(1000000))
		for _, item := range movies {
			if new_movie.Id == item.Id {
				repeatedId = true
				break
			}
		}
		// Condition that only happens when an unique movie "id" is created
		if repeatedId == false {
			break
		}
	}
	movies = append(movies, new_movie)

	responseMessage := fmt.Sprintf("Movie with id: %s, was created successfully", movies[len(movies)-1].Id)
	json.NewEncoder(w).Encode(new_movie)
	fmt.Println(responseMessage)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//TODO

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			// Fancy way of deleting the movie by appending previous and next ones
			movies = append(movies[:index], movies[index+1:]...)

			responseMessage := fmt.Sprintf("Movie with id: %s, was deleted successfully", params["id"])
			json.NewEncoder(w).Encode(Message{Message: responseMessage})
			fmt.Println(responseMessage)

			return
		}
	}
	responseMessage := fmt.Sprintf("Movie with id: %s, was NOT found", params["id"])
	json.NewEncoder(w).Encode(Message{Message: responseMessage})
	fmt.Println(responseMessage)
}

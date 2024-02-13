package main

import (
	
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
)

type movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}
type Director struct{
Firstname string `json:"firstname"`
Lastname string `json:"Lastname"`



}
var movies[] movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item := range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			 break
		}
		
	}
	json.NewEncoder(w).Encode(movies)

}
func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)

	for _,item:=range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	}

}
func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var Movie movie
	_=json.NewDecoder(r.Body).Decode(&Movie)
	Movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies=append(movies, Movie)
	json.NewEncoder(w).Encode(Movie)




}

func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)

	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			var Movie movie
			_=json.NewDecoder(r.Body).Decode(&Movie)
			Movie.ID=params["id"]
			movies =append(movies, Movie)
			json.NewEncoder(w).Encode(Movie)

		}
	}

}


func main(){
	r :=mux.NewRouter()
	movies=append(movies,movie{ID:"1",Isbn:"4322",Title:"Movie one",Director:&Director{Firstname:"John",Lastname:"Doe"}})
	movies=append(movies,movie{ID:"2",Isbn:"43222",Title:"Movie two",Director:&Director{Firstname:"Raj",Lastname:"Mehta"}})
	
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("DELETE")

	fmt.Printf("Starting port at 8000")
	log.Fatal(http.ListenAndServe(":8000",r))



}

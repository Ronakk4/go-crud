package main
import(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	title string `json:"title`
	Director *Director `json:"director"`

}
type Director struct{
Firstname string `json:"firstname"`
Lastname string `json:"Lastname"`



}
var Movie[] movie

func main(){
	r :=mux.NewRouter()
	r.HandleFunc("/movies,",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMoviesByID).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("DELETE")

	fmt.Printf("Starting port at 8000")
	log.Fatal(http.ListenAndServe(":8000",r))



}

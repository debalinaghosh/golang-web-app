package main
import (
	"net/http"
	"log"
  "github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)	
	log.Fatal(http.ListenAndServe(":8000", r));
}

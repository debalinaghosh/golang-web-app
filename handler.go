package main
import (
	"net/http"
)
 func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello First Go App!\n"))	
}

package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Expose-Headers", "*")
}

func fib(n int) int{
	if n == 1 || n == 0 {
		return 1
	}
	return fib(n-1) + fib(n -2)
}

func getFib(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	//get stuff from endpoint body
	//TODO we should be checking we didn't get back junk from this header
	iteration, _ := strconv.Atoi(r.Header.Get("iteration"))
	if iteration > 0 {
			res := fib(iteration)
			fmt.Fprintf(w, strconv.Itoa(res))
		}else {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
}

func handleRequests(){
	myRouter :=mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/fib/v1/iterate", getFib).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	handleRequests()
}
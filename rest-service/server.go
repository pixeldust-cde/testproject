package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"strings"
)

jobs := make(chan int, 100)
results := make(chan int, 100)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Expose-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
    (*w).Header().Set("Access-Control-Allow-Headers", "*")
}


func fib(n int) int{
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n -2)
}

func worker(){
	for n:= range jobs {
		results <- fib(n)
	}
}

type responseJSON struct {
	Number int `json:"result"`
}

func getFib(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	iteration, _ := strconv.Atoi(r.Header.Get("iteration"))
	if iteration >= 0 &&  {

			res := fib(iteration)
			fmt.Fprintf(w, strconv.Itoa(res))
		
		}else {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
}

func handleRequests(){
	myRouter :=mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/fib/v1/iterate", getFib)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}



func main(){
	go worker()
	go worker()
	go worker()
	go worker()
	handleRequests()
}

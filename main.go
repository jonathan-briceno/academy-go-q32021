package main

import (
	"fmt"
	"net/http"

	controllers "academy-go-q32021/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/read-csv/get-cat-fact/{id}", ReadCsv).Methods("GET")
	router.HandleFunc("/read-csv/get-new-cat-fact/", CallExternalApi).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	controllers.HealthCheck(w, r)
}

func ReadCsv(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	controllers.ReadCsv(w, r, params["id"])
}

func CallExternalApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Call External API")
	controllers.GetNewFactFromAPI(w, r)
}

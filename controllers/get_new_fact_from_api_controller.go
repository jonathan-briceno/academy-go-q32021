package controllers

import (
	"academy-go-q32021/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetNewFactFromAPI(w http.ResponseWriter, r *http.Request) {

	cf, err := services.CallExternalApi()

	if err == nil {
		fmt.Println("Now Saving on CSV")
		services.SaveRowOnCsv(cf)
	} else {

		w.WriteHeader(http.StatusInternalServerError)
	}

	factJson, _ := json.Marshal(cf)
	w.Write(factJson)
}

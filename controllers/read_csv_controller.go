package controllers

import (
	"academy-go-q32021/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadCsv(w http.ResponseWriter, r *http.Request, id string) {

	cf, err := repositories.ReadCsv(id)

	fmt.Println("controller reading err", err)
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	cfJson, _ := json.Marshal(cf)
	w.Write(cfJson)
}

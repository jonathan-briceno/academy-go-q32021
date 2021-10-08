package services

import (
	"academy-go-q32021/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallExternalApi() (entities.CatFact, error) {

	fmt.Println("Call External API")

	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	fmt.Println(string(body))              // convert to string before print

	var catfact entities.CatFact
	if err := json.Unmarshal(body, &catfact); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return catfact, err

}

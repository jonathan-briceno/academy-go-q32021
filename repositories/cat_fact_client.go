package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"academy-go-q32021/entities"
)

const (
	BaseUrlV1 = "https://catfact.ninja/fact"
)

func callCatAPI() entities.CatFact {
	resp, err := http.Get(BaseUrlV1)
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

	fmt.Println(catfact)
	return catfact
}

package services

import (
	"academy-go-q32021/entities"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func SaveRowOnCsv(cf entities.CatFact) {

	fmt.Println("Saving on CSV")
	f, err := os.OpenFile("resources/cat_facts.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	csvLines, _ := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	rows := 0
	for range csvLines {
		rows = rows + 1
	}
	rows = rows + 1

	fmt.Println("rows ", rows)
	w := csv.NewWriter(f)
	records := [][]string{
		{strconv.Itoa(rows), string(cf.Fact)},
	}
	for _, empRow := range records {
		error1 := w.Write(empRow)
		fmt.Println("error? ", error1)
	}
	w.Flush()
}

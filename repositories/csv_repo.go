package repositories

import (
	"academy-go-q32021/entities"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type csvRepo interface {
	userExists(string) bool
}

func ReadCsv(id string) (entities.CatFactRow, error) {
	// Open the file

	fmt.Println("reading file...")

	target, targetError := strconv.Atoi(id)
	if targetError != nil {
		fmt.Println("target error")
		return notFound(), errors.New("Fact id must be a valid number")
	}
	csvfile, err := os.Open("resources/cat_facts.csv")
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
		return notFound(), errors.New("An error happened opening the CSV file")
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			return notFound(), nil
		}
		if err != nil {
			fmt.Println("An error has ocurred", err)
			return notFound(), errors.New("An unexpected error happened while processing the search")
		}

		rid, errId := strconv.Atoi(record[0])
		if errId == nil {
			if rid == target {
				return entities.CatFactRow{Id: rid, Fact: record[1]}, nil
			}
		} else {

			return notFound(), errors.New("An error happened reading an Id from the CSV file")
		}
	}
}

func notFound() entities.CatFactRow {
	return entities.CatFactRow{Id: 0, Fact: "Not Found"}
}

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

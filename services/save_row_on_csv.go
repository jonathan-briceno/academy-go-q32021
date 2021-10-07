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
	//f, err := os.OpenFile("resources/cat_facts.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	f, err := os.OpenFile("resources/cat_facts.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	///
	// Parse the file
	csvLines, _ := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	rows := 0
	for range csvLines {
		rows = rows + 1
	}
	rows = rows + 1
	//rows := len(r)
	// Iterate through the records

	///
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

	/*	for i := 0; i < 10; i++ {
			w.Write([]string{"a", "b", "c"})
		}
		w.Flush()
	*/
	/*for _, empRow := range records {
		error1 := w.Write(empRow)
		fmt.Println("error? ", error1)
	}*/

	/*for i := 0; i < 10; i++ {
		w.Write([]string{"a", "b", "c"})
	}*/
	//w.Flush()

	/*if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened CSV file")

		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			fmt.Println(err)
		}

		l := len(csvLines)
	    if len(column) < l {
	        l = len(column)
	    }
	    for i := 0; i < l; i++ {
	        lines[i] = append(lines[i], column[i])
	    }
		/*rows := 0
		for range csvLines {
			rows = rows + 1
		}
		defer csvFile.Close()
		fmt.Println("rows ", rows)*/
	/*svFile, _ := os.Create("resources/cat_factseree.csv")
	csvWriter := csv.NewWriter(csvFile)*/
	//cfr := entities.CatFactRow{Id: rows, Fact: cf.Fact}

	/*records := [][]string{
		{"100", cf.Fact},
	}

	for _, empRow := range records {
		error1 := csvWriter.Write(empRow)
		fmt.Println("error? ", error1)
	}

	defer csvWriter.Flush()*/

}

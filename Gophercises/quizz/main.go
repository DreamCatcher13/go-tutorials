package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func readCsv(filePath string) [][]string { // returns slice of slices
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close() // will close file at the end of a function

	csvReader := csv.NewReader(f)       // create a reader that reads from "f"
	records, err := csvReader.ReadAll() // read all field into slice of strings
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	word := flag.String("w", "hello", "write a word and i will say it back")
	flag.Parse()

	//file := "problems.csv"
	//r := readCsv(file)
	fmt.Printf("word: %v", *word)
}

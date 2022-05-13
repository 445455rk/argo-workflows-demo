package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main(){
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}

	csvReader := csv.NewReader(inFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV ", err)
	}
	inFile.Close()
	os.Rename(os.Args[1],os.Args[1] + "_old")
	outFile, err := os.OpenFile(os.Args[1],os.O_CREATE,0755)
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer outFile.Close()
	csvWriter := csv.NewWriter(outFile)
	err = csvWriter.WriteAll(records[0:5])
	if err != nil {
		log.Fatal("Unable to write CSV ", err)
	}
	csvWriter.Flush()



}

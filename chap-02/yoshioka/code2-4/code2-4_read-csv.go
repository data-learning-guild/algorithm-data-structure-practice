package main

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsv(filepath string) [][]string{

    f, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }

    r := csv.NewReader(f)
    
    record, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    return record

}
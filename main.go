package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lvm/findupes/pkg/csv"
)

const (
	id       csv.Column = "contactID"
	name     csv.Column = "name"
	lastName csv.Column = "name1"
	email    csv.Column = "email"
	zip      csv.Column = "postalZip"
	addr     csv.Column = "address"
)

func main() {
	filename := flag.String("file", "", "CSV Filename")
	flag.Parse()

	if *filename == "" {
		log.Println("[Error] Missing `filename`")
		return
	}

	rows, err := csv.Reader(*filename)
	if err != nil {
		log.Printf("[Error] Reading CSV Failed: %v", err)
		return
	}

	for _, row := range rows {
		fmt.Println(row[id], row[name], row[lastName], row[email])
	}

}

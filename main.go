package main

import (
	"flag"
	"log"

	"github.com/lvm/findupes/internal/findupes"
	"github.com/lvm/findupes/pkg/csv"
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
		log.Printf("[Error] Reading CSV Failed: %v\n", err)
		return
	}

	dupes := make([][]string, 0)
	people := findupes.NewPeople(rows)
	for _, person := range people {
		dupes = append(dupes, person.Compare(people).Export()...)
	}

	if err := csv.Writer("-", dupes); err != nil {
		log.Printf("[Error] failed to write CSV: %v\n", err)
	}

}

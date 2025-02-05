package main

import (
	"flag"
	"fmt"
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
		log.Printf("[Error] Reading CSV Failed: %v", err)
		return
	}

	dupes := make([][]string, 0)
	people := findupes.NewPeople(rows)
	for _, person := range people {
		dupes = append(dupes, person.Compare(people).Export()...)
	}

	for _, dp := range dupes {
		fmt.Println(dp)
	}

}

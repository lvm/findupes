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

	people := findupes.NewPeople(rows)

	for _, p := range people {
		for _, other := range people {
			score := findupes.GetScore(p, other)
			if accuracy := findupes.GetAccuracy(score); accuracy != "" {
				fmt.Println(p.ID, p.FullName(), p.Email, "|", other.ID, other.FullName(), other.Email, "=", accuracy, fmt.Sprintf("(%.2f)", score))
			}
		}
	}

}

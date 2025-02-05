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

type Person struct {
	ID             string
	name, lastName string
	Email          string
	Zip, Address   string
}

func NewPerson(row csv.Row) Person {
	p := new(Person)
	p.ID = row[id]
	p.name = row[name]
	p.lastName = row[lastName]
	p.Email = row[email]
	p.Zip = row[zip]
	p.Address = row[addr]

	return *p
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.name, p.lastName)
}

func getScore(person, other Person) float64 {
	if person.ID == other.ID {
		return 1.0
	}

	values := map[csv.Column]float64{
		name:     0.25,
		lastName: 0.25,
		email:    0.3,
		zip:      0.1,
		addr:     0.1,
	}

	var score float64

	if person.name == other.name {
		score += values[name]
	}

	if person.lastName == other.lastName {
		score += values[lastName]
	}

	if person.Email == other.Email {
		score += values[email]
	}

	if person.Zip == other.Zip {
		score += values[zip]
	}

	if person.Address == other.Address {
		score += values[addr]
	}

	return score
}

func getAccuracy(score float64) string {
	switch {
	case score > 0.25 && score < 0.5:
		return "Low"
	case score > 0.25 && score < 0.5:
		return "Mid"
	case score > 0.25 && score < 0.5:
		return "High"
	default:
		return ""
	}
}

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

	people := make([]Person, 0)
	for _, row := range rows {
		people = append(people, NewPerson(row))
	}

	for _, p := range people {
		for _, other := range people {
			score := getScore(p, other)
			if accuracy := getAccuracy(score); accuracy != "" {
				fmt.Println(p.ID, p.FullName(), p.Email, "|", other.ID, other.FullName(), other.Email, "=", accuracy, fmt.Sprintf("(%.2f)", score))
			}
		}
	}

}

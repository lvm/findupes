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
		p := NewPerson(row)
		fmt.Println(p.ID, p.FullName(), p.Email)
	}

}

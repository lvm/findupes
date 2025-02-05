package findupes

import (
	"fmt"
	"strings"

	"github.com/lvm/findupes/pkg/csv"
)

type (
	Person struct {
		ID             string
		name, lastName string
		Email          string
		Zip, Address   string
	}
	People []Person
)

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

func NewPeople(rows []csv.Row) People {
	people := make(People, len(rows))

	for i, row := range rows {
		people[i] = NewPerson(row)
	}

	return people
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.name, p.lastName)
}

func (p Person) Username() string {
	if p.Email == "" {
		return ""
	}

	parts := strings.Split(p.Email, "@")
	return parts[0]
}

func (p Person) Compare(people People) Results {
	results := make(Results, 0)

	for _, other := range people {
		score := GetScore(p, other)
		if accuracy := GetAccuracy(score); accuracy != nil {
			results = append(results, NewResult(p.ID, other.ID, *accuracy))
		}
	}

	return results
}

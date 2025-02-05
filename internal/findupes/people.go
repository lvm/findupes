package findupes

import (
	"fmt"
	"strings"

	"github.com/lvm/findupes/pkg/csv"
)

type (
	// Each Person represents a row in the data provided
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

/*
As said in io.go, this "method", obtains the Score as the result of the comparison of two Persons.
Each score go from 0.0, to 1.0, no match at all to an exact match. Based on this score, obtains an Accuracy.
*/
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

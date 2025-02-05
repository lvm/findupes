package findupes

import (
	"log"
	"os"

	"github.com/lvm/findupes/pkg/csv"
)

type Importer func(f *os.File) ([]csv.Row, error)
type Exporter func(f *os.File, content [][]string) error

func Process(importer Importer, ifile *os.File, exporter Exporter, efile *os.File) error {
	rows, err := importer(ifile)
	if err != nil {
		log.Printf("[Error] Reading CSV failed: %v\n", err)
		return err
	}

	people := NewPeople(rows)
	dupes := [][]string{{string(source), string(match), string(accuracy)}}
	for _, person := range people {
		dupes = append(dupes, person.Compare(people).Export()...)
	}

	if err := exporter(efile, dupes); err != nil {
		log.Printf("[Error] failed to write CSV: %v\n", err)
	}

	return nil
}

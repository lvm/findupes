package findupes

import (
	"io"
	"log"

	"github.com/lvm/findupes/pkg/csv"
)

type Importer func(r io.Reader) ([]csv.Row, error)
type Exporter func(w io.Writer, content [][]string) error

func Process(importer Importer, rbuff io.Reader, exporter Exporter, wbuff io.Writer) error {
	rows, err := importer(rbuff)
	if err != nil {
		log.Printf("[Error] Reading CSV failed: %v\n", err)
		return err
	}

	people := NewPeople(rows)
	dupes := [][]string{{string(source), string(match), string(accuracy)}}
	for _, person := range people {
		dupes = append(dupes, person.Compare(people).Export()...)
	}

	if err := exporter(wbuff, dupes); err != nil {
		log.Printf("[Error] failed to write CSV: %v\n", err)
	}

	return nil
}

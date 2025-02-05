package findupes

import (
	"io"
	"log"

	"github.com/lvm/findupes/pkg/csv"
)

type Importer func(r io.Reader) ([]csv.Row, error)
type Exporter func(w io.Writer, content [][]string) error

/*
Process takes care of reading an input-file and writing an output-file.
Importer and Exporter, normally use csv.Reader and csv.Writer (see `pkg/csv`), but can be extended to use json.Reader and json.Exporter or any encoding it's prefered.

Initially, the data is being read by an encoding Reader (Importer), then the heavy load is done by each "Person".
Person.Compare(People) compares 1-to-1, so it's O(n^2). I guess it could be rewritten to get closer (if not) O(1), but due to time constraint, wasn't possible.
Person.Compare(People).Export() returns a slice of string slices, each containing three columns:  ContactID Source, ContactID Match, Accuracy
Finally, that result will be sent to an encoding Writer (Exporter).

Only returning an error if any step failed.
*/
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

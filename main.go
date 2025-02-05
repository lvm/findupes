package main

import (
	"flag"
	"log"
	"os"

	"github.com/lvm/findupes/internal/findupes"
	"github.com/lvm/findupes/pkg/csv"
)

func main() {
	var (
		ifile, efile         *os.File
		importCSV, exportCSV *string
		err                  error
	)

	importCSV = flag.String("import", "", "CSV Import")
	exportCSV = flag.String("export", "-", "CSV Export (default: stdout)")
	flag.Parse()

	if *importCSV == "" {
		log.Println("[Error] Missing input CSV file")
		return
	}

	ifile, err = os.Open(*importCSV)
	if err != nil {
		log.Printf("[Error] Opening input CSV file failed: %v\n", err)
		return
	}
	defer ifile.Close()

	if *exportCSV == "-" {
		efile = os.Stdout
	} else {
		efile, err = os.OpenFile(*exportCSV, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Printf("[Error] Opening output CSV file failed: %v\n", err)
		}
		defer efile.Close()
	}

	if err := findupes.Process(csv.Reader, ifile, csv.Writer, efile); err != nil {
		log.Printf("[Error] Something went wrong processing the Contacts: %v\n", err)
		return
	}

}

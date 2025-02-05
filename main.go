package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/lvm/findupes/internal/findupes"
	"github.com/lvm/findupes/pkg/csv"
)

func main() {
	var (
		importCSV, exportCSV *string
		err                  error
	)

	importCSV = flag.String("import", "", "CSV Import (required)")
	exportCSV = flag.String("export", "-", "CSV Export")
	flag.Parse()

	if *importCSV == "" {
		flag.Usage()
		return
	}

	rbuff, err := os.Open(*importCSV)
	if err != nil {
		log.Printf("[Error] Opening input CSV file failed: %v\n", err)
		return
	}
	defer rbuff.Close()

	var wbuff io.Writer
	if *exportCSV == "-" {
		wbuff = os.Stdout
	} else {
		wbuff, err = os.OpenFile(*exportCSV, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Printf("[Error] Opening output CSV file failed: %v\n", err)
			return
		}
		defer wbuff.(*os.File).Close()
	}

	if err := findupes.Process(csv.Reader, rbuff, csv.Writer, wbuff); err != nil {
		log.Printf("[Error] Something went wrong processing the Contacts: %v\n", err)
		return
	}

}

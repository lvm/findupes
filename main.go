package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

type Column string
type Header []Column
type Row map[Column]string

const (
	id       Column = "contactID"
	name     Column = "name"
	lastName Column = "name1"
	email    Column = "email"
	zip      Column = "postalZip"
	addr     Column = "address"
)

var ErrNotEnoughData = errors.New("csv seems to be empty")

func NewHeader(cols ...string) Header {
	header := make(Header, len(cols))

	for i, col := range cols {
		header[i] = Column(col)
	}

	return header
}

func csvReader(filename string) ([]Row, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(lines) < 2 {
		return nil, ErrNotEnoughData
	}

	headers := NewHeader(lines[0]...)

	var rows []Row
	for _, line := range lines[1:] {
		row := make(Row)
		for i, value := range line {
			row[headers[i]] = value
		}
		rows = append(rows, row)
	}

	return rows, nil
}

func main() {
	filename := flag.String("file", "", "CSV Filename")
	flag.Parse()

	if *filename == "" {
		log.Println("[Error] Missing `filename`")
		return
	}

	rows, err := csvReader(*filename)
	if err != nil {
		log.Println("[Error] Reading CSV Failed: %v", err)
		return
	}

	for _, row := range rows {
		fmt.Println(row[id], row[name], row[lastName], row[email])
	}

}

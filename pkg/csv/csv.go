package csv

import (
	"encoding/csv"
	"errors"
	"os"
)

type Column string
type Header []Column
type Row map[Column]string

var (
	ErrNoFilename    = errors.New("no filename")
	ErrNotEnoughData = errors.New("csv seems to be empty")
)

func NewHeader(cols ...string) Header {
	header := make(Header, len(cols))

	for i, col := range cols {
		header[i] = Column(col)
	}

	return header
}

func Reader(filename string) ([]Row, error) {
	if filename == "" {
		return nil, ErrNoFilename
	}

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

func Writer(filename string, content [][]string) error {
	if filename == "" {
		return ErrNoFilename
	}

	var file *os.File

	if filename == "-" {
		file = os.Stdout
	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(content); err != nil {
		return err
	}

	writer.Flush()

	return writer.Error()
}

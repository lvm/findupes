package csv

import (
	"encoding/csv"
	"errors"
	"io"
)

type (
	Column string
	Header []Column
	Row    map[Column]string
)

var (
	ErrNoFile        = errors.New("missing file")
	ErrNotEnoughData = errors.New("csv seems to be empty")
)

func NewHeader(cols ...string) Header {
	header := make(Header, len(cols))

	for i, col := range cols {
		header[i] = Column(col)
	}

	return header
}

func Reader(r io.Reader) ([]Row, error) {
	if r == nil {
		return nil, ErrNoFile
	}

	reader := csv.NewReader(r)
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

func Writer(w io.Writer, content [][]string) error {
	if w == nil {
		return ErrNoFile
	}

	writer := csv.NewWriter(w)
	if err := writer.WriteAll(content); err != nil {
		return err
	}

	writer.Flush()
	return writer.Error()
}

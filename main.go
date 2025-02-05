package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "", "CSV Filename")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return
	}

	if len(rows) < 2 {
		return
	}

	headers := rows[0]

	var lines []map[string]string
	for _, row := range rows[1:] {
		line := make(map[string]string)
		for i, value := range row {
			line[headers[i]] = value
		}
		lines = append(lines, line)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

}

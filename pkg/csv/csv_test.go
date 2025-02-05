package csv

import (
	"bytes"
	"fmt"
	"testing"
)

type mockWriter struct{}

func (m *mockWriter) Write(_ []byte) (int, error) {
	return -1, fmt.Errorf("oops")
}

func TestReader(t *testing.T) {

	t.Run("Valid CSV - Reader", func(t *testing.T) {
		data := `contactID,name,name1,email,postalZip,address
1,Ciara,French,mollis.lectus.pede@outlook.net,39746,449-6990 Tellus. Rd.
2,Charles,Pacheco,nulla.eget@protonmail.couk,76837,Ap #312-8611 Lacus. Ave`
		file := bytes.NewBufferString(data)

		rows, err := Reader(file)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(rows) != 2 {
			t.Fatalf("expected 2 rows, got %d", len(rows))
		}

		expected := []Row{
			{"contactID": "1", "name": "Ciara", "name1": "French", "email": "mollis.lectus.pede@outlook.net", "postalZip": "39746", "address": "449-6990 Tellus. Rd."},
			{"contactID": "2", "name": "Charles", "name1": "Pacheco", "email": "nulla.eget@protonmail.couk", "postalZip": "76837", "address": "Ap #312-8611 Lacus. Ave"},
		}

		for i, row := range expected {
			for k, v := range row {
				if rows[i][k] != v {
					t.Errorf("expected %s, got %s for column %s", v, rows[i][k], k)
				}
			}
		}
	})

	t.Run("Empty CSV", func(t *testing.T) {
		if _, err := Reader(bytes.NewBufferString("")); err != ErrNotEnoughData {
			t.Fatalf("expected error %v, got %v", ErrNotEnoughData, err)
		}
	})

	t.Run("Just Header", func(t *testing.T) {
		data := `contactID,name,name1,email,postalZip,address`
		file := bytes.NewBufferString(data)
		_, err := Reader(file)
		if err != ErrNotEnoughData {
			t.Fatalf("expected error %v, got %v", ErrNotEnoughData, err)
		}
	})
}

func TestWriter(t *testing.T) {

	t.Run("Valid CSV - Writer", func(t *testing.T) {
		var buf bytes.Buffer
		content := [][]string{
			{"contactID", "name", "name1", "email", "postalZip", "address"},
			{"1", "Ciara", "French", "mollis.lectus.pede@outlook.net", "39746", "449-6990 Tellus. Rd."},
			{"2", "Charles", "Pacheco", "nulla.eget@protonmail.couk", "76837", "Ap #312-8611 Lacus. Ave"},
		}

		if err := Writer(&buf, content); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := "contactID,name,name1,email,postalZip,address\n1,Ciara,French,mollis.lectus.pede@outlook.net,39746,449-6990 Tellus. Rd.\n2,Charles,Pacheco,nulla.eget@protonmail.couk,76837,Ap #312-8611 Lacus. Ave\n"
		if buf.String() != expected {
			t.Errorf("expected:\n%s\nbut got:\n%s", expected, buf.String())
		}
	})

	t.Run("No file", func(t *testing.T) {
		if err := Writer(nil, [][]string{}); err != ErrNoFile {
			t.Fatalf("expected error %v, got %v", ErrNoFile, err)
		}
	})

	t.Run("Writer Error", func(t *testing.T) {
		writer := &mockWriter{}
		if err := Writer(writer, [][]string{{"contactID", "name", "name1", "email", "postalZip", "address"}}); err == nil {
			t.Fatal("expected error when writing, got nil")
		}
	})
}

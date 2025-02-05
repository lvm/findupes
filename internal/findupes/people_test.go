package findupes

import (
	"reflect"
	"testing"

	"github.com/lvm/findupes/pkg/csv"
)

func TestPerson(t *testing.T) {

	t.Run("NewPerson", func(t *testing.T) {
		rows := []csv.Row{
			{"contactID": "1", "name": "Ciara", "name1": "French", "email": "mollis.lectus.pede@outlook.net", "postalZip": "39746", "address": "449-6990 Tellus. Rd."},
			{"contactID": "2", "name": "Charles", "name1": "Pacheco", "email": "nulla.eget@protonmail.couk", "postalZip": "76837", "address": "Ap #312-8611 Lacus. Ave"},
		}

		p := NewPerson(rows[0])
		if p.ID != "1" {
			t.Errorf("expected %v got %v", 1, p.ID)
		}

		ps := NewPeople(rows)
		if ps[0].ID != "1" && ps[1].ID != "2" {
			t.Errorf("expected %v and %v, got %v and %v", 1, ps[0].ID, 2, ps[1].ID)
		}

	})

	t.Run("Fullname", func(t *testing.T) {
		p := NewPerson(csv.Row{"contactID": "1", "name": "Ciara", "name1": "French", "email": "mollis.lectus.pede@outlook.net", "postalZip": "39746", "address": "449-6990 Tellus. Rd."})
		expected := "Ciara French"
		if p.FullName() != expected {
			t.Errorf("expected %v got %v", expected, p.FullName())
		}
	})

	t.Run("Username", func(t *testing.T) {
		p := NewPerson(csv.Row{"contactID": "1", "name": "Ciara", "name1": "French", "email": "mollis.lectus.pede@outlook.net", "postalZip": "39746", "address": "449-6990 Tellus. Rd."})
		expected := "mollis.lectus.pede"
		if p.Username() != expected {
			t.Errorf("expected %v got %v", expected, p.Username())
		}

		p_ := NewPerson(csv.Row{"contactID": "1", "name": "Ciara", "name1": "French", "email": "", "postalZip": "39746", "address": "449-6990 Tellus. Rd."})
		if p_.Username() != "" {
			t.Errorf("expected %v got %v", "", p_.Username())
		}
	})

	t.Run("Compare", func(t *testing.T) {
		rows := []csv.Row{
			{"contactID": "1", "name": "Ciara", "name1": "French", "email": "mollis.lectus.pede@outlook.net", "postalZip": "39746", "address": "449-6990 Tellus. Rd."},
			{"contactID": "2", "name": "Charles", "name1": "Pacheco", "email": "nulla.eget@protonmail.couk", "postalZip": "76837", "address": "Ap #312-8611 Lacus. Ave"},
			{"contactID": "3", "name": "Charles", "name1": "Lopez", "email": "eget@protonmail.couk", "postalZip": "55555", "address": "312 Lacus. Ave"},
			{"contactID": "4", "name": "Ciara", "name1": "French", "email": "mollis.lectus.pede@outlook.com", "postalZip": "39746", "address": "449-6990 Tellus. Road"},
		}

		ps := NewPeople(rows)
		if ps[0].ID != "1" && ps[1].ID != "2" {
			t.Errorf("expected %v and %v, got %v and %v", 1, ps[0].ID, 2, ps[1].ID)
		}

		result := ps[0].Compare(ps)
		expected := Results{
			{SourceID: "1", MatchID: "1", Accuracy: Match},
			{SourceID: "1", MatchID: "4", Accuracy: Mid},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %+v got %+v", expected, result)
		}

	})

}

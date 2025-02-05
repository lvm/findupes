package findupes

import (
	"testing"

	"github.com/lvm/findupes/pkg/csv"
)

func TestScore(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		s1 := Score(10)
		s2 := Score(5)

		s1.Sum(s2)
		if s1 != 15 {
			t.Errorf("expected 15, got %v", s1)
		}
	})

	t.Run("Half", func(t *testing.T) {
		s := Score(10)
		result := s.Half()
		if result != 5 {
			t.Errorf("expected 5, got %v", result)
		}
	})

	t.Run("Lte", func(t *testing.T) {
		s1 := Score(5)
		s2 := Score(10)

		if !s1.Lte(s2) {
			t.Errorf("expected %v to be less than or equal to %v", s1, s2)
		}
	})

	t.Run("Gte", func(t *testing.T) {
		s1 := Score(10)
		s2 := Score(5)

		if !s1.Gte(s2) {
			t.Errorf("expected %v to be greater than or equal to %v", s1, s2)
		}
	})

	t.Run("Eq", func(t *testing.T) {
		s1 := Score(5)
		s2 := Score(5)
		s3 := Score(10)

		if !s1.Eq(s2) {
			t.Errorf("expected %v to be equal to %v", s1, s2)
		}

		if s1.Eq(s3) {
			t.Errorf("expected %v to not be equal to %v", s1, s2)
		}
	})

	t.Run("GetScore", func(t *testing.T) {
		rows := []csv.Row{
			{"contactID": "1", "name": "Charles", "name1": "Pacheco", "email": "nulla.eget@gmail.couk", "postalZip": "39746", "address": "449-6990 Tellus. Rd."},
			{"contactID": "2", "name": "Charles", "name1": "Pacheco", "email": "nulla.eget@protonmail.couk", "postalZip": "76837", "address": "Ap #312-8611 Lacus. Ave"},
		}

		ppl := NewPeople(rows)
		score := GetScore(ppl[0], ppl[1])
		expected := Score(0.45)

		if score != expected {
			t.Errorf("expected %v got %v", expected, score)
		}
	})
}

func TestAccuracy(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		expected := "Test"
		acc := Accuracy(expected)
		if acc.String() != expected {
			t.Errorf("expected '%v', got %v", expected, acc.String())
		}
	})

	t.Run("GetAccuracy", func(t *testing.T) {
		res := GetAccuracy(Score(0.1))
		if res != nil {
			t.Errorf("expected `%+v`, got `%+v`", nil, res)
		}

		var lo, mid, hi, mat Accuracy = Low, Mid, Hi, Match
		testCases := []struct {
			score    Score
			expected Accuracy
		}{
			{score: Score(0.3), expected: lo},
			{score: Score(0.6), expected: mid},
			{score: Score(0.8), expected: hi},
			{score: Score(1.0), expected: mat},
		}

		for _, tc := range testCases {
			res := GetAccuracy(tc.score)
			if *res != tc.expected {
				t.Errorf("expected `%+v`, got `%+v`", tc.expected, res)
			}
		}

	})
}

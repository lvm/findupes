package findupes

import "github.com/lvm/findupes/pkg/csv"

type (
	Score    float64
	Accuracy string
)

func (s *Score) Sum(v Score) {
	*s += v
}

func (s Score) Half() Score {
	return s / 2
}

func (s Score) Lte(v Score) bool {
	return s <= v
}

func (s Score) Gte(v Score) bool {
	return s >= v
}

func (s Score) Eq(v Score) bool {
	return s == v
}

func (a Accuracy) String() string {
	return string(a)
}

func GetScore(person, other Person) Score {
	if person.ID == other.ID {
		return 1.0
	}

	values := map[csv.Column]Score{
		fullName: Score(0.45),
		email:    Score(0.3),
		zip:      Score(0.1),
		addr:     Score(0.1),
		bonus:    Score(0.05),
	}

	var score Score

	if person.FullName() == other.FullName() {
		score.Sum(values[fullName])
	}

	if person.Email != "" && other.Email == "" {
		if person.Email == other.Email {
			score.Sum(values[email])
		} else if person.Username() == other.Username() {
			score.Sum(values[email].Half())
		}

		// extra score if name and email matches.
		if person.Email == other.Email && person.FullName() == other.FullName() {
			score.Sum(values[bonus])
		}
	}

	if person.Zip == other.Zip {
		score.Sum(values[zip])
	}

	if person.Address == other.Address {
		score.Sum(values[addr])
	}

	return score
}

func GetAccuracy(score Score) *Accuracy {
	var acc Accuracy

	switch {
	case score.Gte(0.25) && score.Lte(0.49):
		acc = Low
	case score.Gte(0.5) && score.Lte(0.74):
		acc = Mid
	case score.Gte(0.75) && score.Lte(0.9):
		acc = Hi
	case score.Eq(1.0):
		acc = Match
	default:
		return nil
	}

	return &acc
}

package findupes

import "github.com/lvm/findupes/pkg/csv"

func GetScore(person, other Person) float64 {
	if person.ID == other.ID {
		return 1.0
	}

	values := map[csv.Column]float64{
		name:     0.25,
		lastName: 0.25,
		email:    0.3,
		zip:      0.1,
		addr:     0.1,
	}

	var score float64

	if person.name == other.name {
		score += values[name]
	}

	if person.lastName == other.lastName {
		score += values[lastName]
	}

	if person.Email == other.Email {
		score += values[email]
	}

	if person.Zip == other.Zip {
		score += values[zip]
	}

	if person.Address == other.Address {
		score += values[addr]
	}

	return score
}

func GetAccuracy(score float64) string {
	switch {
	case score > 0.25 && score < 0.5:
		return "Low"
	case score > 0.25 && score < 0.5:
		return "Mid"
	case score > 0.25 && score < 0.5:
		return "High"
	default:
		return ""
	}
}

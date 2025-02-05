package findupes

import "github.com/lvm/findupes/pkg/csv"

const (
	id       csv.Column = "contactID"
	name     csv.Column = "name"
	lastName csv.Column = "name1"
	email    csv.Column = "email"
	zip      csv.Column = "postalZip"
	addr     csv.Column = "address"

	// non csv columns, but used for scoring
	fullName csv.Column = "fullName"
	bonus    csv.Column = "bonus"
)

const (
	Low   Accuracy = "Low"
	Mid   Accuracy = "Mid"
	Hi    Accuracy = "High"
	Match Accuracy = "Match"
)

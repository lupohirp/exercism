package isbnverifier

import (
    "strconv"  
    "strings"
)



func IsValidISBN(isbn string) bool {

	isbnWithoutHyphens := strings.ReplaceAll(isbn, "-", "")

	if len(isbnWithoutHyphens) != 10 {
		return false
	}

	multiplier := 10
	sumOfDigits := 0

	for i, v := range isbnWithoutHyphens {

		var digit int
		var err error

		digit, err = strconv.Atoi(string(v))

		if err != nil {
			if i == len(isbnWithoutHyphens)-1 {
				if string(v) != "X" {
					return false
				} else {
					digit = 10
				}
			} else {
				return false
			}
		}

		sumOfDigits += (multiplier * digit)
		multiplier--
	}

	return sumOfDigits%11 == 0

}


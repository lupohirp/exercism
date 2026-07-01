package luhn

import ( 
	"strings"
    "strconv"
)

func Valid(id string) bool {


    trimmedString := strings.ReplaceAll(id, " ", "")

	if len(trimmedString) <= 1 {
		return false
	}



	sumOfDgts := 0

	shouldDouble := false

	for i := len(trimmedString) - 1; i >= 0; i-- {

		digit, err := strconv.Atoi(string(trimmedString[i]))
        if err != nil{
            return false
        }
		if !shouldDouble {
			sumOfDgts += digit
		} else {
			doubling := digit * 2
			if doubling > 9 {
				doubling = doubling - 9
			}
			sumOfDgts += doubling
		}

		shouldDouble = !shouldDouble

	}

	return sumOfDgts%10 == 0
}

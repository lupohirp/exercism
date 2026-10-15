package phonenumber

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"
)


func Number(phoneNumber string) (string, error) {
	reg, _ := regexp.Compile("[^0-9]")
	cleanedNumber := reg.ReplaceAllString(phoneNumber, "")
	cleanedNumber = strings.TrimSpace(cleanedNumber)

	if len(cleanedNumber) == 11 {
		if cleanedNumber[0:1] == "1" {
			cleanedNumber = cleanedNumber[1:]
		} else {
			return "", fmt.Errorf("Invalid phone number")
		}
	} else if len(cleanedNumber) != 10 {
		return "", fmt.Errorf("Invalid phone number")
	}

	firstDigit := cleanedNumber[0:1]
	firstDigitAsInt, _ := strconv.Atoi(firstDigit)

	if firstDigitAsInt < 2 || firstDigitAsInt > 9 {
		return "", fmt.Errorf("Invalid phone number")
	}

	thirdDigit := cleanedNumber[3:4]

	thirdDigitAsInt, _ := strconv.Atoi(thirdDigit)

	if thirdDigitAsInt < 2 || thirdDigitAsInt > 9 {
		return "", fmt.Errorf("Invalid phone number")
	}

	return cleanedNumber, nil

}


func AreaCode(phoneNumber string) (string, error) {

	reg, _ := regexp.Compile("[^0-9]")
	cleanedNumber := reg.ReplaceAllString(phoneNumber, "")
	cleanedNumber = strings.TrimSpace(cleanedNumber)

	if len(cleanedNumber) == 11 {
		if cleanedNumber[0:1] == "1" {
			cleanedNumber = cleanedNumber[1:]
		} else {
			return "", fmt.Errorf("Invalid phone number")
		}
	} else if len(cleanedNumber) != 10 {
		return "", fmt.Errorf("Invalid phone number")
	}

	firstDigit := cleanedNumber[0:1]
	firstDigitAsInt, _ := strconv.Atoi(firstDigit)

	if firstDigitAsInt < 2 || firstDigitAsInt > 9 {
		return "", fmt.Errorf("Invalid phone number")
	}

    thirdDigit := cleanedNumber[3:4]

	thirdDigitAsInt, _ := strconv.Atoi(thirdDigit)

	if thirdDigitAsInt < 2 || thirdDigitAsInt > 9 {
		return "", fmt.Errorf("Invalid phone number")
	}

	return cleanedNumber[0:3], nil

}

func Format(phoneNumber string) (string, error) {
	reg, _ := regexp.Compile("[^0-9]")
	cleanedNumber := reg.ReplaceAllString(phoneNumber, "")
	cleanedNumber = strings.TrimSpace(cleanedNumber)

	if len(cleanedNumber) == 11 {
		if cleanedNumber[0:1] == "1" {
			cleanedNumber = cleanedNumber[1:]
		} else {
			return "", fmt.Errorf("Invalid phone number")
		}
	} else if len(cleanedNumber) != 10 {
		return "", fmt.Errorf("Invalid phone number")
	}

    
		firstDigit := cleanedNumber[0:1]
	firstDigitAsInt, _ := strconv.Atoi(firstDigit)

	if firstDigitAsInt < 2 || firstDigitAsInt > 9 {
		return "", fmt.Errorf("Invalid phone number")
	}

    thirdDigit := cleanedNumber[3:4]

	thirdDigitAsInt, _ := strconv.Atoi(thirdDigit)

	if thirdDigitAsInt < 2 || thirdDigitAsInt > 9 {
		return "", fmt.Errorf("Invalid phone number")
	}

	areaCode := "(" + cleanedNumber[0:3] + ") "
	nxx := cleanedNumber[3:6]
	number := cleanedNumber[6:10]

	return areaCode + nxx + "-" + number, nil

}


package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b){
        return 0,errors.New("Distances not equals")
    }

    distance := 0

    for i, char := range a {

        if(string(b[i]) != string(char)){
            distance++
        }
    }

    return distance,nil
}

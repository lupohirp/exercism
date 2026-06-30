package collatzconjecture

import "errors"
      

func CollatzConjecture(n int) (int, error) {

		if n <= 0  {
            return 0, errors.New("error")
        }

    	numberOfSteps := 0

    	if n == 1 {
            return numberOfSteps,nil
        }
    
		for {

            if n % 2 == 0{
                n = n/2
            } else if n % 2 != 0 {
                n = n*3 +1
            }

            numberOfSteps++

            if(n == 1){
                break
            }
        }

    return numberOfSteps, nil

}

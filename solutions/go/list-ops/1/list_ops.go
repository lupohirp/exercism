package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}

	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	reverse := make([]int, len(s))

	count := len(s) - 1

	for _, v := range s {

		reverse[count] = v
		count--
	}

	acc := initial
	for _, v := range reverse {
		acc = fn(v, acc)
	}

	return acc

}

func (s IntList) Filter(fn func(int) bool) IntList {

	var filteredS = make(IntList, len(s))
	var i int = 0

	for _, v := range s {
		if fn(v) {
			filteredS[i] = v
			i++
		}
	}

	filteredS = filteredS[:i]

	return filteredS
}

func (s IntList) Length() int {
	var count int = 0

	for range s {
		count++
	}

	return count
}

func (s IntList) Map(fn func(int) int) IntList {

	result := make(IntList, len(s))
	i := 0

	for _, v := range s {
		result[i] = fn(v)
		i++
	}

	return result
}

func (s IntList) Reverse() IntList {

	reverse := make([]int, len(s))

	count := len(s) - 1

	for _, v := range s {

		reverse[count] = v
		count--
	}

	return reverse

}

func (s IntList) Append(lst IntList) IntList {

	toReturn := make([]int, len(s)+len(lst))

	var last int = 0

	for _, v := range s {
		toReturn[last] = v
		last++
	}

	for _, v := range lst {
		toReturn[last] = v
		last++
	}

	return toReturn

}

func (s IntList) Concat(lists []IntList) IntList {

	maxCapacity := 0
	for _, list := range lists {
		for range list {
			maxCapacity++
		}
	}

	toReturn := make([]int, maxCapacity)
	var last int = 0

	for _, list := range lists {
		for _, v := range list {
			toReturn[last] = v
			last++
		}
	}

	return toReturn
}

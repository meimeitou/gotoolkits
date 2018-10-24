package util

import "strconv"

func StrToIntSlice(a []string) (error, []int, int) {
	ints := make([]int, len(a))
	for index, str := range a {
		i, err := strconv.Atoi(str)
		if err != nil {
			return err, ints, index
		}
		ints[index] = i
	}

	return nil, ints, len(ints) - 1
}

func StrToInt64Slice(a []string) (error, []int64, int) {
	ints := make([]int64, len(a))
	for index, str := range a {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err, ints, index
		}
		ints[index] = i
	}

	return nil, ints, len(ints) - 1
}

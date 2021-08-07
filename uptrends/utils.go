package uptrends

import (
	"strconv"
)

func Bool(input bool) *bool {
	return &input
}

func Int(input int) *int {
	return &input
}

func Int32(input int32) *int32 {
	return &input
}

func Int64(input int64) *int64 {
	return &input
}

func Float(input float64) *float64 {
	return &input
}

func String(input string) *string {
	return &input
}

func SliceInterfaceToSliceInt32(input []interface{}) *[]int32 {
	ret := make([]int32, len(input))

	for i, x := range input {
		switch x := x.(type) {
		case int:
			ret[i] = int32(x)
		case string:
			s, _ := strconv.Atoi(x)
			ret[i] = int32(s)
		}
	}

	return &ret
}

func SliceInterfaceToSliceString(input []interface{}) *[]string {
	ret := make([]string, len(input))

	for i, x := range input {
		ret[i] = string(x.(string))
	}

	return &ret
}

func SliceInt32ToSliceString(input []int32) *[]string {
	ret := make([]string, len(input))

	for i, x := range input {
		s := strconv.Itoa(int(x))
		ret[i] = s
	}

	return &ret
} 
package uptrends

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wasfree/uptrends-go-sdk"
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

// SliceInterfaceToSliceInt32 convert []interface{} to *[]int32
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

// SliceInterfaceToSliceString convert []interface{} to *[]string
func SliceInterfaceToSliceString(input []interface{}) *[]string {
	ret := make([]string, len(input))

	for i, x := range input {
		ret[i] = string(x.(string))
	}

	return &ret
}

// SliceInt32ToSliceString convert []int32 to []string
func SliceInt32ToSliceString(input []int32) []string {
	ret := make([]string, len(input))

	for i, x := range input {
		s := strconv.Itoa(int(x))
		ret[i] = s
	}

	return ret
}

// reverseMap takes a map[string]string and reverse key and value order
func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

// MergeSchema will merge schemaA into schemaB and return a schema consist of both schemas
func MergeSchema(schemaA, schemaB map[string]*schema.Schema) map[string]*schema.Schema {
	for k, v := range schemaA {
		schemaB[k] = v
	}

	return schemaB
}

// SliceInterfaceToSliceRequestHeader converts values from []interface{} to []uptrends.RequestHeader
func SliceInterfaceToSliceRequestHeader(input []interface{}) *[]uptrends.RequestHeader {
	headers := []uptrends.RequestHeader{}
	for _, v := range input {
		newHeader := uptrends.NewRequestHeader()
		header := v.(map[string]interface{})
		newHeader.SetName(header["name"].(string))
		newHeader.SetValue(header["value"].(string))
		headers = append(headers, *newHeader)
	}
	return &headers
}

// SliceRequestHeaderToSliceInterface converts values from []uptrends.RequestHeader to []interface{}
func SliceRequestHeaderToSliceInterface(input []uptrends.RequestHeader) []interface{} {
	headers := []interface{}{}
	for _, v := range input {
		newHeader := make(map[string]interface{})
		newHeader["name"] = v.GetName()
		newHeader["value"] = v.GetValue()
		headers = append(headers, newHeader)
	}
	return headers
}

// SliceInterfaceToSlicePatternMatch converts values from []interface{} to []uptrends.PatternMatch
func SliceInterfaceToSlicePatternMatch(input []interface{}) *[]uptrends.PatternMatch {
	patterns := []uptrends.PatternMatch{}
	for _, v := range input {
		pattern := v.(map[string]interface{})
		newPattern := uptrends.NewPatternMatch(pattern["is_positive"].(bool))
		newPattern.SetPattern(pattern["pattern"].(string))
		patterns = append(patterns, *newPattern)
	}
	return &patterns
}

func SlicePatternMatchToSliceInterface(input []uptrends.PatternMatch) []interface{} {
	patterns := []interface{}{}
	for _, v := range input {
		newPattern := make(map[string]interface{})
		newPattern["pattern"] = v.GetPattern()
		newPattern["is_positive"] = v.IsPositive
		patterns = append(patterns, newPattern)
	}
	return patterns
}

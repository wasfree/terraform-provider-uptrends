package uptrends

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/wasfree/uptrends-go-sdk"
)

func TestSliceInterfaceToSliceInt32(t *testing.T) {
	slice := []interface{}{123, 456, 789, "1234"}
	result := SliceInterfaceToSliceInt32(slice)

	final := []int32{123, 456, 789, 1234}
	assert.Equal(t, &final, result)
}

func TestMergeSchema(t *testing.T) {
	schemaA := map[string]*schema.Schema{
		"example1": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Just a example.",
		},
	}
	schemaB := map[string]*schema.Schema{
		"example2": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Just a example.",
		},
	}
	schemaC := map[string]*schema.Schema{
		"example3": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Just a example.",
		},
	}

	final := MergeSchema(schemaA, schemaB, schemaC)

	assert.Len(t, final, 3)
	assert.Contains(t, final, "example1")
	assert.Contains(t, final, "example2")
	assert.Contains(t, final, "example3")
}

func TestRequestHeaderHelpers(t *testing.T) {
	requestHeaderInterfaceTestData := []interface{}{
		map[string]interface{}{
			"name":  "Content-Type",
			"value": "text/html; charset=UTF-8",
		},
	}
	requestHeaderTestData := []uptrends.RequestHeader{
		{
			Name:  String("Content-Type"),
			Value: String("text/html; charset=UTF-8"),
		},
	}

	t.Run("Test expandRequestHeader", func(t *testing.T) {
		headers := expandRequestHeader(requestHeaderInterfaceTestData)
		assert.Equal(t, requestHeaderTestData, headers)
	})
	t.Run("Test flattenRequestHeader", func(t *testing.T) {
		headers := flattenRequestHeader(&requestHeaderTestData)
		assert.Equal(t, requestHeaderInterfaceTestData, headers)
	})
}

func TestPatternMatchHelper(t *testing.T) {
	patternMatchInterfaceTestData := []interface{}{
		map[string]interface{}{
			"pattern":     "somepattern",
			"is_positive": true,
		},
	}
	patternMatchTestData := []uptrends.PatternMatch{
		{
			Pattern:    String("somepattern"),
			IsPositive: true,
		},
	}

	t.Run("Test expandPatternMatch", func(t *testing.T) {
		patterns := expandPatternMatch(patternMatchInterfaceTestData)
		assert.Equal(t, patternMatchTestData, patterns)
	})
	t.Run("Test flattenPatternMatch", func(t *testing.T) {
		patterns := flattenPatternMatch(&patternMatchTestData)
		assert.Equal(t, patternMatchInterfaceTestData, patterns)
	})
}

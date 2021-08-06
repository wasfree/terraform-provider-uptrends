package uptrends

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSliceInterfaceToSliceInt32(t *testing.T) {
	slice := []interface{}{123, 456, 789}
	result := SliceInterfaceToSliceInt32(slice)

	final := []int32{123, 456, 789}
	assert.Equal(t, &final, result)
}
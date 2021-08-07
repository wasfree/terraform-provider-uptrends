package uptrends

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceInterfaceToSliceInt32(t *testing.T) {
	slice := []interface{}{123, 456, 789, "1234"}
	result := SliceInterfaceToSliceInt32(slice)

	final := []int32{123, 456, 789, 1234}
	assert.Equal(t, &final, result)
}

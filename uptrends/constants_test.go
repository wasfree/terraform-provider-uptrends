package uptrends

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckpointID(t *testing.T) {
	ckpt1 := CheckpointID("London")
	assert.Equal(t, "1", ckpt1, "Expect to resolve ID from checkpoint London")

	ckpt2 := CheckpointID("1")
	assert.Equal(t, "1", ckpt2, "Expect to get ID if ID was provided")

	ckpt3 := CheckpointID("New York")
	assert.Equal(t, "8", ckpt3, "Expect to get ID from checkpoint name with space")
}

func TestDeduplicateCheckpointIDs(t *testing.T) {
	ckpts := []interface{}{"New York", "New York", "8", "Vancouver", "Gloucester", "10", "10"}
	ids := DeduplicateCheckpointIDs(ckpts)
	assert.Equal(t, []int32{8, 14, 17, 10}, *ids, "expect to return list without duplicate ids")
}

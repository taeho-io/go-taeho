package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachineID(t *testing.T) {
	id, err := machineID()
	assert.NotZero(t, id)
	assert.Nil(t, err)
}

func TestNew(t *testing.T) {
	tid := New()
	id, err := tid.Generate()
	assert.NotZero(t, id)
	assert.Nil(t, err)
}

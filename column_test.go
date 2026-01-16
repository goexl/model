package model_test

import (
	"testing"

	"github.com/goexl/model"
	"github.com/stretchr/testify/assert"
)

func TestColumn(t *testing.T) {
	assert.Equal(t, "Created", model.Created)
	assert.Equal(t, "Updated", model.Updated)
	assert.Equal(t, "Deleted", model.Deleted)
}

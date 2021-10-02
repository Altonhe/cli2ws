package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	assert.NotNil(t, err())
}

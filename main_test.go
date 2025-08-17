package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	expected := "Hello, Coffee!"
	result := HelloCoffee()
	assert.Equal(t, expected, result, "Error: HelloCoffee()")
}

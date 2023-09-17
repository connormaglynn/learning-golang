package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	assert.Equal(t, "hello world!", Hello("hello world!"))
}

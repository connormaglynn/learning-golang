package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("returns Hello, Connor when Connor is passed", func(t *testing.T) {
		assert.Equal(t, "Hello, Connor", Hello("Connor"))
	})

	t.Run("returns Hello, World when empty string is passed", func(t *testing.T) {
		assert.Equal(t, "Hello, World", Hello(""))
	})
}

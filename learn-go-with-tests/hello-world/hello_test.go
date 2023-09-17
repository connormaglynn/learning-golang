package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("returns 'Hello, World' when empty string and no language is passed", func(t *testing.T) {
		assert.Equal(t, "Hello, World", Hello("", ""))
	})

	t.Run("returns 'Hello, Connor' when 'Connor' and no language is passed", func(t *testing.T) {
		assert.Equal(t, "Hello, Connor", Hello("Connor", ""))
	})

	t.Run("returns 'Hola, World' when empty string and 'Spanish'", func(t *testing.T) {
		assert.Equal(t, "Hola, World", Hello("", "Spanish"))
	})

}

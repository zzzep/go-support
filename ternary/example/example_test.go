package example

import (
	"github.com/stretchr/testify/assert"
	"github.com/zzzep/go-support/ternary"
	"testing"
)

func TestIfTernary(t *testing.T) {
	t.Run("TestIfTernary", func(t *testing.T) {
		res := ternary.If(true, 1, 2)
		assert.Equal(t, 1, res)
	})
	t.Run("TestIfTernary2", func(t *testing.T) {
		firstValue := 3
		secondValue := 7
		res := ternary.If(firstValue < secondValue, 1, 2)
		assert.Equal(t, 1, res)
	})
	t.Run("ternary Run function", func(t *testing.T) {
		myFunc := ternary.If(true, func() int {
			return 1
		}, func() int {
			return 2
		})
		assert.Equal(t, 1, myFunc())
	})
	t.Run("false", func(t *testing.T) {
		assert.Equal(t, 2, ternary.If(false, 1, 2))
	})
}

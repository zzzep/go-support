package setter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetIfNotEmptyString(t *testing.T) {
	t.Run("should set value when not empty", func(t *testing.T) {
		var target string
		SetIfNotEmptyString(&target, "hello")
		assert.Equal(t, "hello", target)
	})

	t.Run("should not set value when empty", func(t *testing.T) {
		target := "existing"
		SetIfNotEmptyString(&target, "")
		assert.Equal(t, "existing", target)
	})

	t.Run("should not set value when empty on empty target", func(t *testing.T) {
		var target string
		SetIfNotEmptyString(&target, "")
		assert.Equal(t, "", target)
	})
}

func TestSetIfNotEmptySlice(t *testing.T) {
	t.Run("should set value when not empty", func(t *testing.T) {
		var target []string
		SetIfNotEmptySlice(&target, []string{"a", "b"})
		assert.Equal(t, []string{"a", "b"}, target)
	})

	t.Run("should not set value when empty", func(t *testing.T) {
		target := []string{"existing"}
		SetIfNotEmptySlice(&target, []string{})
		assert.Equal(t, []string{"existing"}, target)
	})

	t.Run("should not set value when nil", func(t *testing.T) {
		target := []string{"existing"}
		SetIfNotEmptySlice(&target, nil)
		assert.Equal(t, []string{"existing"}, target)
	})
}

func TestSetIfNotEmptyMap(t *testing.T) {
	t.Run("should set value when not empty", func(t *testing.T) {
		var target map[string]int
		SetIfNotEmptyMap(&target, map[string]int{"a": 1})
		assert.Equal(t, map[string]int{"a": 1}, target)
	})

	t.Run("should not set value when empty", func(t *testing.T) {
		target := map[string]int{"existing": 1}
		SetIfNotEmptyMap(&target, map[string]int{})
		assert.Equal(t, map[string]int{"existing": 1}, target)
	})

	t.Run("should not set value when nil", func(t *testing.T) {
		target := map[string]int{"existing": 1}
		SetIfNotEmptyMap(&target, nil)
		assert.Equal(t, map[string]int{"existing": 1}, target)
	})
}

func TestSetIfNotEmpty(t *testing.T) {
	t.Run("should set int value when not zero", func(t *testing.T) {
		var target int
		SetIfNotEmpty(&target, 42)
		assert.Equal(t, 42, target)
	})

	t.Run("should not set int value when zero", func(t *testing.T) {
		target := 10
		SetIfNotEmpty(&target, 0)
		assert.Equal(t, 10, target)
	})

	t.Run("should set string value when not empty", func(t *testing.T) {
		var target string
		SetIfNotEmpty(&target, "hello")
		assert.Equal(t, "hello", target)
	})

	t.Run("should not set string value when empty", func(t *testing.T) {
		target := "existing"
		SetIfNotEmpty(&target, "")
		assert.Equal(t, "existing", target)
	})

	t.Run("should set bool value when true", func(t *testing.T) {
		var target bool
		SetIfNotEmpty(&target, true)
		assert.Equal(t, true, target)
	})

	t.Run("should not set bool value when false", func(t *testing.T) {
		target := true
		SetIfNotEmpty(&target, false)
		assert.Equal(t, true, target)
	})
}

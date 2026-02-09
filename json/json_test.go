package json

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestRedoUnmarshal(t *testing.T) {
	t.Run("Success with map[string]interface{}", func(t *testing.T) {
		from := map[string]interface{}{"key": "value"}
		got, err := RedoUnmarshal[map[string]any](from)
		assert.NoError(t, err)
		assert.Equal(t, map[string]any{"key": "value"}, got)
	})
	
	t.Run("Success with map[string]any", func(t *testing.T) {
		from := map[string]any{"key": "value"}
		got, err := RedoUnmarshal[map[string]any](from)
		assert.NoError(t, err)
		assert.Equal(t, map[string]any{"key": "value"}, got)
	})
	
	t.Run("Success with struct", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		from := Person{Name: "John", Age: 30}
		got, err := RedoUnmarshal[Person](from)
		assert.NoError(t, err)
		assert.Equal(t, Person{Name: "John", Age: 30}, got)
	})
	
	t.Run("Success with array", func(t *testing.T) {
		from := []int{1, 2, 3}
		got, err := RedoUnmarshal[[]int](from)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, got)
	})
	
	t.Run("Error with unmarshalable type", func(t *testing.T) {
		// Channel cannot be marshaled
		from := make(chan int)
		got, err := RedoUnmarshal[map[string]any](from)
		assert.Error(t, err)
		assert.Empty(t, got)
	})
	
	t.Run("Error with function type", func(t *testing.T) {
		// Function cannot be marshaled
		from := func() {}
		got, err := RedoUnmarshal[map[string]any](from)
		assert.Error(t, err)
		assert.Empty(t, got)
	})
}

func TestUnmarshal(t *testing.T) {
	t.Run("Success with []byte", func(t *testing.T) {
		data := []byte(`{"key": "value"}`)
		got, err := Unmarshal[map[string]any](data)
		assert.NoError(t, err)
		assert.Equal(t, map[string]any{"key": "value"}, got)
	})
	
	t.Run("Success with string", func(t *testing.T) {
		data := `{"key": "value"}`
		got, err := Unmarshal[map[string]any](data)
		assert.NoError(t, err)
		assert.Equal(t, map[string]any{"key": "value"}, got)
	})
	
	t.Run("Error with type mismatch", func(t *testing.T) {
		data := []byte(`{"key": "erro"}`)
		got, err := Unmarshal[map[string]int](data)
		assert.Error(t, err)
		assert.Equal(t, map[string]int{"key": 0}, got)
	})
	
	t.Run("Invalid JSON with []byte", func(t *testing.T) {
		var want map[string]any
		data := []byte(`invalid-json`)
		got, err := Unmarshal[map[string]any](data)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
	
	t.Run("Invalid JSON with string", func(t *testing.T) {
		var want map[string]any
		data := "invalid-json"
		got, err := Unmarshal[map[string]any](data)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
	
	t.Run("io.Reader success", func(t *testing.T) {
		want := map[string]any{
			"key": "value",
		}
		var data io.Reader
		data = bytes.NewReader([]byte(`{"key": "value"}`))
		got, err := Unmarshal[map[string]any](data)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
	
	t.Run("io.Reader with invalid JSON", func(t *testing.T) {
		var want map[string]any
		data := bytes.NewReader([]byte(`invalid-json`))
		got, err := Unmarshal[map[string]any](data)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
	
	t.Run("Unsupported type", func(t *testing.T) {
		var want map[string]any
		data := 123 // int is not supported
		got, err := Unmarshal[map[string]any](data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type")
		assert.Equal(t, want, got)
	})
	
	t.Run("Unsupported type with float", func(t *testing.T) {
		var want map[string]any
		data := 123.45 // float is not supported
		got, err := Unmarshal[map[string]any](data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type")
		assert.Equal(t, want, got)
	})
	
	t.Run("Unsupported type with bool", func(t *testing.T) {
		var want map[string]any
		data := true // bool is not supported
		got, err := Unmarshal[map[string]any](data)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type")
		assert.Equal(t, want, got)
	})
	
	t.Run("Complex struct with string", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		data := `{"name": "John", "age": 30}`
		got, err := Unmarshal[Person](data)
		assert.NoError(t, err)
		assert.Equal(t, Person{Name: "John", Age: 30}, got)
	})
	
	t.Run("Array with string", func(t *testing.T) {
		data := `[1, 2, 3]`
		got, err := Unmarshal[[]int](data)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, got)
	})
}

func TestMarshal(t *testing.T) {
	t.Run("Success with map", func(t *testing.T) {
		data := map[string]any{"key": "value"}
		got, err := Marshal[map[string]any](data)
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"key":"value"}`), got)
	})
	
	t.Run("Success with struct", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		data := Person{Name: "John", Age: 30}
		got, err := Marshal[Person](data)
		assert.NoError(t, err)
		assert.Contains(t, string(got), "John")
		assert.Contains(t, string(got), "30")
	})
	
	t.Run("Success with array", func(t *testing.T) {
		data := []int{1, 2, 3}
		got, err := Marshal[[]int](data)
		assert.NoError(t, err)
		assert.Equal(t, []byte(`[1,2,3]`), got)
	})
	
	t.Run("Error with channel", func(t *testing.T) {
		// Channel cannot be marshaled
		data := make(chan int)
		got, err := Marshal[chan int](data)
		assert.Error(t, err)
		assert.Nil(t, got)
	})
	
	t.Run("Error with function", func(t *testing.T) {
		// Function cannot be marshaled
		data := func() {}
		got, err := Marshal[func()](data)
		assert.Error(t, err)
		assert.Nil(t, got)
	})
}

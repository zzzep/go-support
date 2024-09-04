package json

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedoUnmarshal(t *testing.T) {
	type args struct {
		from any
		to   any
	}
	type testCase struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testCase{
		{name: "TestRedoUnmarshal", args: args{from: map[string]interface{}{"key": "value"}, to: map[string]any{}}, wantErr: assert.NoError},
		{name: "TestRedoUnmarshal", args: args{from: map[string]any{"key": "value"}, to: map[string]interface{}{}}, wantErr: assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := RedoUnmarshal[map[string]any](tt.args.from)
			if (err != nil) != tt.wantErr(t, err, fmt.Sprintf("RedoUnmarshal(%v, %v)", tt.args.from, tt.args.to)) {
				return
			}
			assert.Equalf(t, tt.args.to, data, "RedoUnmarshal(%v, %v)", tt.args.from, tt.args.to)
		})
	}
}

func TestUnmarshal(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		data := []byte(`{"key": "value"}`)
		got, err := Unmarshal[map[string]any](data)
		assert.NoError(t, err)
		assert.Equal(t, map[string]any{"key": "value"}, got)
	})
	t.Run("Error", func(t *testing.T) {
		data := []byte(`{"key": "erro"}`)
		got, err := Unmarshal[map[string]int](data)
		assert.Error(t, err)
		assert.Equal(t, map[string]int{"key": 0}, got)
	})
	t.Run("invalid json", func(t *testing.T) {
		var want map[string]any
		data := []byte(`invalid-json`)
		got, err := Unmarshal[map[string]any](data)
		assert.Error(t, err)
		assert.Equal(t, want, got)
	})
}

func TestMarshal(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		data := map[string]any{"key": "value"}
		got, err := Marshal[map[string]any](data)
		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"key":"value"}`), got)
	})
}

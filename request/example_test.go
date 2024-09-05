package request

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type exRes struct {
	Data string `json:"data"`
}

func TestHttp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()
	t.Run("GET", func(t *testing.T) {
		res, err := SimpleGet[exRes](context.Background(), ts.URL, nil)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})

	t.Run("POST", func(t *testing.T) {
		res, err := SimplePost[exRes](context.Background(), ts.URL, nil, map[string]string{
			"test": "test",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})

	t.Run("PUT", func(t *testing.T) {
		res, err := SimplePut[exRes](context.Background(), ts.URL, nil, map[string]string{
			"test": "test",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})

	t.Run("PATCH", func(t *testing.T) {
		res, err := SimplePatch[exRes](context.Background(), ts.URL, nil, map[string]string{
			"test": "test",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})

	t.Run("DELETE", func(t *testing.T) {
		res, err := SimpleDelete[exRes](context.Background(), ts.URL, nil, map[string]string{
			"test": "test",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})

	t.Run("OPTIONS", func(t *testing.T) {
		res, err := SimpleOptions[exRes](context.Background(), ts.URL, nil, map[string]string{
			"test": "test",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})

	t.Run("Request", func(t *testing.T) {
		res, err := SimpleRequest[exRes](context.Background(), ts.URL, "GET", nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})

	t.Run("Request with body", func(t *testing.T) {
		res, err := SimpleRequest[exRes](context.Background(), ts.URL, "POST", nil, map[string]string{
			"test": "test",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "test", res.Data)
	})
}

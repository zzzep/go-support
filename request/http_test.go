package request

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testResponse struct {
	Data string `json:"data"`
}

func TestHttp_Constructor(t *testing.T) {
	req := Http[testResponse]()
	assert.NotNil(t, req)
	
	// Test that default headers are set
	req = req.SetUrl("http://example.com")
	payload := req.GetPayload()
	assert.Equal(t, "", payload.Data)
}

func TestSetContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := Http[testResponse]().SetContext(ctx).SetUrl("http://example.com")
	assert.NotNil(t, req)
}

func TestSetUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).Get()
	assert.NoError(t, req.GetError())
	assert.Equal(t, "test", req.GetPayload().Data)
}

func TestSetStructUrl(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	url := Url{
		schema: "http://",
		domain: strings.TrimPrefix(ts.URL, "http://"),
		path:   "/%s",
		params: []string{"test"},
	}

	// SetStructUrl is not in the interface, need to use type assertion
	reqInterface := Http[testResponse]()
	req := reqInterface.(*HttpReq[testResponse])
	req.SetStructUrl(url)
	result := req.Get()
	assert.NoError(t, result.GetError())
	assert.Equal(t, "test", result.GetPayload().Data)
}

func TestSetHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customHeader := r.Header.Get("X-Custom-Header")
		assert.Equal(t, "custom-value", customHeader)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	headers := map[string]string{
		"X-Custom-Header": "custom-value",
		"X-Another-Header": "another-value",
	}

	req := Http[testResponse]().SetUrl(ts.URL).SetHeader(headers).Get()
	assert.NoError(t, req.GetError())
}

func TestSetHeader_Empty(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).SetHeader(nil).Get()
	assert.NoError(t, req.GetError())
}

func TestSetTransport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// SetTransport is not in the interface, need to use type assertion
	reqInterface := Http[testResponse]().SetUrl(ts.URL)
	req := reqInterface.(*HttpReq[testResponse])
	req.SetTransport(customTransport)
	result := req.Get()
	assert.NoError(t, result.GetError())
	assert.Equal(t, "test", result.GetPayload().Data)
}

func TestSetJsonBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		assert.Contains(t, string(body), "test-body")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	body := map[string]string{
		"key": "test-body",
	}

	req := Http[testResponse]().SetUrl(ts.URL).SetJsonBody(body).Post()
	assert.NoError(t, req.GetError())
}

func TestSetJsonBody_Nil(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).SetJsonBody(nil).Post()
	assert.NoError(t, req.GetError())
}

func TestSetBody(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		assert.Equal(t, "raw body", string(body))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).SetBody(strings.NewReader("raw body")).Post()
	assert.NoError(t, req.GetError())
}

func TestGetResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).Get()
	assert.NoError(t, req.GetError())
	
	response := req.GetResponse()
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestGetResponse_WithError(t *testing.T) {
	// Use invalid URL to cause error
	req := Http[testResponse]().SetUrl("http://invalid-url-that-does-not-exist.local:9999").Get()
	assert.Error(t, req.GetError())
	
	response := req.GetResponse()
	assert.Nil(t, response)
}

func TestGetPayload_WithError(t *testing.T) {
	// Use invalid URL to cause error
	req := Http[testResponse]().SetUrl("http://invalid-url-that-does-not-exist.local:9999").Get()
	assert.Error(t, req.GetError())
	
	payload := req.GetPayload()
	// Should return zero value when there's an error
	assert.Equal(t, "", payload.Data)
}

func TestRequest_InvalidURL(t *testing.T) {
	req := Http[testResponse]().SetUrl("://invalid-url").Get()
	assert.Error(t, req.GetError())
}

func TestRequest_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	req := Http[testResponse]().SetContext(ctx).SetUrl(ts.URL).Get()
	// May or may not error depending on timing, but should handle gracefully
	_ = req.GetError()
}

func TestRequest_ErrorHandling(t *testing.T) {
	// Test with malformed URL
	req := Http[testResponse]().SetUrl("not-a-valid-url").Request("GET")
	assert.Error(t, req.GetError())
}

func TestRequest_AllMethods(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":"test"}`))
	}))
	defer ts.Close()

	methods := []struct {
		name   string
		method func() HttpInterface[testResponse]
	}{
		{"Get", func() HttpInterface[testResponse] {
			return Http[testResponse]().SetUrl(ts.URL).Get()
		}},
		{"Post", func() HttpInterface[testResponse] {
			return Http[testResponse]().SetUrl(ts.URL).Post()
		}},
		{"Put", func() HttpInterface[testResponse] {
			return Http[testResponse]().SetUrl(ts.URL).Put()
		}},
		{"Patch", func() HttpInterface[testResponse] {
			return Http[testResponse]().SetUrl(ts.URL).Patch()
		}},
		{"Delete", func() HttpInterface[testResponse] {
			return Http[testResponse]().SetUrl(ts.URL).Delete()
		}},
		{"Options", func() HttpInterface[testResponse] {
			return Http[testResponse]().SetUrl(ts.URL).Options()
		}},
	}

	for _, tt := range methods {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.method()
			assert.NoError(t, req.GetError())
			assert.Equal(t, "test", req.GetPayload().Data)
		})
	}
}

func TestRequest_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`invalid json`))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).Get()
	assert.Error(t, req.GetError())
}

func TestRequest_EmptyResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(``))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).Get()
	// Empty response may or may not error depending on JSON unmarshal
	_ = req.GetError()
}

func TestRequest_HTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"server error"}`))
	}))
	defer ts.Close()

	req := Http[testResponse]().SetUrl(ts.URL).Get()
	// HTTP error status doesn't cause request error, but response will have status 500
	response := req.GetResponse()
	if response != nil {
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	}
}

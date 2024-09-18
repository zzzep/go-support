package request

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/zzzep/go-support/json"
	"io"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ContentTypeJson = "application/json"
)

var (
	defaultHeaders = http.Header{
		ContentType: {ContentTypeJson},
	}
	defaultTransport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
)

type HttpInterface[T any] interface {
	SetUrl(url string) HttpInterface[T]
	SetContext(ctx context.Context) HttpInterface[T]
	SetHeader(headers map[string]string) HttpInterface[T]
	SetJsonBody(body any) HttpInterface[T]
	SetBody(body io.Reader) HttpInterface[T]
	Request(method string) HttpInterface[T]
	Get() HttpInterface[T]
	Post() HttpInterface[T]
	Patch() HttpInterface[T]
	Put() HttpInterface[T]
	Delete() HttpInterface[T]
	Options() HttpInterface[T]

	GetPayload() T
	GetError() error
	GetResponse() *http.Response
}

type HttpReq[T any] struct {
	ctx       context.Context
	url       string
	headers   http.Header
	body      io.Reader
	transport *http.Transport
	request   *http.Request

	response *http.Response
	payload  T
	err      error
}

type Url struct {
	schema string
	domain string
	path   string
	params []string
}

func Http[T any]() HttpInterface[T] {
	return &HttpReq[T]{
		ctx:       context.Background(),
		headers:   defaultHeaders,
		transport: defaultTransport,
	}
}

func (hr *HttpReq[T]) SetContext(ctx context.Context) HttpInterface[T] {
	hr.ctx = ctx
	return hr
}

func (hr *HttpReq[T]) SetUrl(url string) HttpInterface[T] {
	hr.url = url
	return hr
}

func (hr *HttpReq[T]) SetStructUrl(url Url) HttpInterface[T] {
	hr.url = url.schema + url.domain + fmt.Sprintf(url.path, url.params)
	return hr
}

func (hr *HttpReq[T]) SetHeader(headers map[string]string) HttpInterface[T] {
	for keyHeader, valueHeader := range headers {
		hr.headers.Add(keyHeader, valueHeader)
	}
	return hr
}

func (hr *HttpReq[T]) SetTransport(transport *http.Transport) HttpInterface[T] {
	hr.transport = transport
	return hr
}

func (hr *HttpReq[T]) SetJsonBody(body any) HttpInterface[T] {
	if body == nil {
		return hr
	}
	bodyJson, _ := json.Marshal(body)
	return hr.SetBody(bytes.NewBuffer(bodyJson))
}

func (hr *HttpReq[T]) SetBody(body io.Reader) HttpInterface[T] {
	hr.body = body
	return hr
}

func (hr *HttpReq[T]) Request(method string) HttpInterface[T] {
	hr.request, hr.err = http.NewRequestWithContext(hr.ctx, method, hr.url, hr.body)
	if hr.err != nil {
		return hr
	}
	hr.err = hr.do(hr.request)
	if hr.err != nil {
		return hr
	}
	defer hr.close()

	resBody, _ := io.ReadAll(hr.response.Body)

	hr.payload, hr.err = json.Unmarshal[T](resBody)

	return hr
}

func (hr *HttpReq[T]) do(req *http.Request) (err error) {
	req.Header = hr.headers

	client := &http.Client{Transport: hr.transport}

	hr.response, err = client.Do(req)

	return err
}

func (hr *HttpReq[T]) close() {
	if hr.response != nil {
		_ = hr.response.Body.Close()
	}
}

func (hr *HttpReq[T]) Get() HttpInterface[T] {
	return hr.Request(http.MethodGet)
}

func (hr *HttpReq[T]) Post() HttpInterface[T] {
	return hr.Request(http.MethodPost)
}

func (hr *HttpReq[T]) Put() HttpInterface[T] {
	return hr.Request(http.MethodPut)
}

func (hr *HttpReq[T]) Delete() HttpInterface[T] {
	return hr.Request(http.MethodDelete)
}

func (hr *HttpReq[T]) Patch() HttpInterface[T] {
	return hr.Request(http.MethodPatch)
}

func (hr *HttpReq[T]) Options() HttpInterface[T] {
	return hr.Request(http.MethodOptions)
}

func (hr *HttpReq[T]) GetPayload() T {
	if hr.err != nil {
		return *new(T)
	}
	return hr.payload
}

func (hr *HttpReq[T]) GetError() error {
	return hr.err
}

func (hr *HttpReq[T]) GetResponse() *http.Response {
	if hr.err != nil {
		return nil
	}
	return hr.response
}

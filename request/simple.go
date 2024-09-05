package request

import "context"

func baseRequest[T any](ctx context.Context, url string, headers map[string]string, body any) HttpInterface[T] {
	return Http[T]().SetContext(ctx).SetUrl(url).SetHeader(headers).SetJsonBody(body)
}

func SimpleRequest[T any](ctx context.Context, url, method string, headers map[string]string, body any) (T, error) {
	req := baseRequest[T](ctx, url, headers, body).Request(method)

	return req.GetPayload(), req.GetError()
}

func SimpleGet[T any](ctx context.Context, url string, headers map[string]string) (T, error) {
	req := baseRequest[T](ctx, url, headers, nil).Get()
	return req.GetPayload(), req.GetError()
}

func SimplePost[T any](ctx context.Context, url string, headers map[string]string, body any) (T, error) {
	req := baseRequest[T](ctx, url, headers, body).Post()
	return req.GetPayload(), req.GetError()
}

func SimplePut[T any](ctx context.Context, url string, headers map[string]string, body any) (T, error) {
	req := baseRequest[T](ctx, url, headers, body).Put()
	return req.GetPayload(), req.GetError()
}

func SimplePatch[T any](ctx context.Context, url string, headers map[string]string, body any) (T, error) {
	req := baseRequest[T](ctx, url, headers, body).Patch()
	return req.GetPayload(), req.GetError()
}

func SimpleDelete[T any](ctx context.Context, url string, headers map[string]string, body any) (T, error) {
	req := baseRequest[T](ctx, url, headers, body).Delete()
	return req.GetPayload(), req.GetError()
}

func SimpleOptions[T any](ctx context.Context, url string, headers map[string]string, body any) (T, error) {
	req := baseRequest[T](ctx, url, headers, body).Options()
	return req.GetPayload(), req.GetError()
}

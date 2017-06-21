package utils

import (
	"net/http"
	"errors"
)

type Request struct {
	Path    string
	Data    interface{}
	Headers map[string]string
}

func (r *Request) Get() (*http.Response, error) {
	if r.Data == nil {
		return http.Get(r.Path)
	} else if query_str, ok := r.Data.(map[string]string); ok {
		httpClient := &http.Client{}
		request, err := http.NewRequest(http.MethodGet, r.Path, nil)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		for key, val := range r.Headers {
			request.Header.Add(key, val)
		}
		q := request.URL.Query()
		for key, val := range query_str {
			q.Add(key, val)
		}
		request.URL.RawQuery = q.Encode()
		return httpClient.Do(request)
	}
	return nil, errors.New("Can't handle GET with data")
}

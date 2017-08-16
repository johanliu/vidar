package vidar

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter
	status        int
	contentType   string
	contentLength int
}

func (res *Response) Write(data []byte) (int, error) {
	if res.status == 0 {
		res.status = http.StatusOK
	}

	size, err := res.ResponseWriter.Write(data)
	res.contentLength += size

	return size, err
}

func (res *Response) SetContentType(value string) error {
	return res.SetHeader("Content-Type", value)
}

func (res *Response) SetStatus(code int) error {
	res.ResponseWriter.WriteHeader(code)
	return nil
}

func (res *Response) SetHeader(key string, value string) error {
	if len(value) == 0 {
		res.ResponseWriter.Header().Del(key)
	} else {
		res.ResponseWriter.Header().Set(key, value)
	}

	return nil
}

func (res *Response) Size() (int, error) {
	return res.contentLength, nil
}

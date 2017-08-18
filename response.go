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

func (res *Response) SetStatus(code int) {
	res.ResponseWriter.WriteHeader(code)
}

func (res *Response) SetHeader(key string, value string) {
	res.ResponseWriter.Header().Set(key, value)
}

func (res *Response) RemoveHeader(key string) {
	res.ResponseWriter.Header().Del(key)
}

func (res *Response) Size() (int, error) {
	return res.contentLength, nil
}

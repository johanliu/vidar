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

func (res *Response) SetBody(data []byte) (int, error) {
	if res.status == 0 {
		res.status = http.StatusOK
	}

	size, err := res.ResponseWriter.Write(data)
	res.contentLength += size

	return size, err
}

func (res *Response) SetStatus(code int) {
	res.status = code
	res.ResponseWriter.WriteHeader(code)
}

func (res *Response) Status() int {
	return res.status
}

func (res *Response) SetHeader(key string, value string) {
	res.ResponseWriter.Header().Set(key, value)
}

func (res *Response) AddHeader(key string, value string) {
	res.ResponseWriter.Header().Add(key, value)
}

func (res *Response) DelHeader(key string) {
	res.ResponseWriter.Header().Del(key)
}

func (res *Response) Size() (int, error) {
	return res.contentLength, nil
}

func (res *Response) Flush() {
	res.ResponseWriter.(http.Flusher).Flush()
}

func (res *Response) Hijack() {
	res.ResponseWriter.(http.Hijacker).Hijack()
}

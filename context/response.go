package context

import "net/http"

type Response struct {
	http.ResponseWriter
	status        int
	contentType   string
	contentLength int
}

func (response *Response) Write(data []byte) (int, error) {
	if response.status == 0 {
		response.status = http.StatusOK
	}

	size, err := response.ResponseWriter.Write(data)
	response.contentLength += size

	return size, err
}

func (response *Response) SetContentType(value string) error {
	return response.SetHeader("Content-Type", value)
}

func (response *Response) SetStatus(code int) error {
	response.ResponseWriter.WriteHeader(code)

	return nil
}

func (response *Response) SetHeader(key string, value string) error {
	if len(value) == 0 {
		response.ResponseWriter.Header().Del(key)
	} else {
		response.ResponseWriter.Header().Set(key, value)
	}

	return nil
}

func (response *Response) GetHeader(key string) string {
	return response.ResponseWriter.Header().Get(key)
}

func (response *Response) Size() (int, error) {
	return response.contentLength, nil
}

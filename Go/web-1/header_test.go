package web_1

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)

	fmt.Println(string(body))
}

func ResponeHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Agis Satria Mandala")
	fmt.Fprint(writer, "OK")
}

func TestResponeHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	recoder := httptest.NewRecorder()

	ResponeHeader(recoder, request)

	respone := recoder.Result()
	body, _ := io.ReadAll(respone.Body)

	fmt.Println(string(body))

	poweredBy := recoder.Header().Get("X-Powered-By")
	fmt.Println(poweredBy)
}

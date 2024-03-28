package web_1

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponeCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponeCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?name=Agis", nil)
	recorder := httptest.NewRecorder()

	ResponeCode(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)

	fmt.Println(respone.StatusCode)
	fmt.Println(respone.Status)
	fmt.Println(string(body))
}

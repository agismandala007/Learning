package web_1

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(writter http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writter, "Hello")
	} else {
		fmt.Fprintf(writter, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?name=Agis", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)

	fmt.Println(string(body))
}

func MultipleParameter(writter http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("firstname")
	lastName := request.URL.Query().Get("lastname")

	fmt.Fprintf(writter, "%s %s", firstName, lastName)
}

func TestMultipleParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?firstname=Agis&lastname=Mandala", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)

	fmt.Println(string(body))
}

func MultipleValueQuery(writer http.ResponseWriter, request *http.Request) {
	var query url.Values = request.URL.Query()
	var names []string = query["name"]

	fmt.Fprintln(writer, strings.Join(names, ", "))
}

func TestMultipleValueQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/hello?name=Agis&name=Mandala", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQuery(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)

	fmt.Println(string(body))
}

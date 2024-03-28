package web_1

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostFormValue("first_name")
	lastName := request.PostFormValue("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Agis&last_name=Mandala")
	request := httptest.NewRequest(http.MethodPost, "http://localhost/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)

	fmt.Println(string(body))
}

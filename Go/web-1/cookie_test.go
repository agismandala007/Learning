package web_1

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = ("X-ASM-Name")
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success Create Cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-ASM-Name")

	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestAddCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/?name=Agis", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCokie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-ASM-Name"
	cookie.Value = "Agis Satria Mandala"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)
	fmt.Println(string(body))
}

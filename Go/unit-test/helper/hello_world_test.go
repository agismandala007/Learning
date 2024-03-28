package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Agis")

	if result != "Hello Agis" {
		t.Error()
	}
}

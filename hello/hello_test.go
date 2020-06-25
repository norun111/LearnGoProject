package hello

import "testing"

func TestGetHello(t *testing.T) {
	result := GetHello("tomoya")
	expect := "Hello tomoya"

	if result != expect {
		t.Error("\nresult: ", result, "\nexpext: ", expect)
	}

	t.Log("test terminated")
}
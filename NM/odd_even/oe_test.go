package odd_even

import "testing"

func TestEvenOperation(t *testing.T) {
	result := EvenOperation(4)
	expect := true

	if result != expect {
		t.Error("\nresult: ", result, "\nexpect: ", expect)
	}

	t.Log("test terminated")
}

package solution

import "testing"

func TestGetMessage(t *testing.T) {

	var expected = string([]byte{72, 101, 108, 108, 111, 32, 240, 159, 151, 186, 239, 184, 143, 32, 33})

	actual := GetMessage()

	if actual != expected {
		t.Errorf("GetMessage() returns: \"%v\". Expected result: \"%v\"", actual, expected)
	}
}


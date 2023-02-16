package addresses

import "testing"

func Test_AddressType(t *testing.T) {
	expected := "Street"
	result := AddressType("Street Banana")

	if expected != result {
		t.Error("Result different from expected")
	}
}

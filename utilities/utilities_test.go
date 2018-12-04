package utilities

import "testing"

func TestStringsToInts(t *testing.T) {
	var input = []string{"0", "3", "0", "1", "-3"}
	var actual = []int{0, 3, 0, 1, -3}
	expected := StringsToInts(input)
	if !testEq(actual, expected) {
		t.Fatalf("failed")
	}
}

func TestIntsToString(t *testing.T) {
	var expected = "0,3,0,1,-3"
	var input = []int{0, 3, 0, 1, -3}
	actual := IntsToString(input)
	if actual != expected {
		t.Fatalf("failed: actual: %s, expected %s", actual, expected)
	}
}

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

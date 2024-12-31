package utils

import "testing"

func TestBlank(t *testing.T) {
	aString := "not blank"

	// Check that Blank() should return false for a non-empty string
	if Blank(aString) == true {
		t.Error("Expected Blank() to return false, but it returned true")
	}

	// Additional test case for an empty string
	emptyString := ""
	if Blank(emptyString) == false {
		t.Error("Expected Blank() to return true for an empty string, but it returned false")
	}

	// Additional test case for a string with only spaces
	spaceString := "    "
	if Blank(spaceString) == false {
		t.Error("Expected Blank() to return true for a string with only spaces, but it returned false")
	}
}

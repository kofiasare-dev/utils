package utils

import (
	"log"
	"testing"
)

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

func TestMapWithIndex(t *testing.T) {
	wa := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	r := MapWithIndex(wa, func(d, i int) int {
		println(d, i)
		return d + i
	})

	log.Print(r)
}

func TestMapSum(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	log.Print(Sum(n))
}

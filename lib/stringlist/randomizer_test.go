package stringlist

import (
	"strings"
	"testing"
)

func assertEqualS(expected, actual string, message string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %#v, but got %#v - %s", expected, actual, message)
	}
}

func assertEqualI(expected, actual int, message string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %#v, but got %#v - %s", expected, actual, message)
	}
}

func TestStuff(t *testing.T) {
	rnd := MakeRandomizer()
	rnd.Append("string 1")
	rnd.Append("string 2")
	rnd.Append("string 3")

	masterListString := strings.Join(rnd.masterList.data, ",")
	assertEqualS("string 1,string 2,string 3", masterListString, "Stringlist value", t)

	assertEqualI(3, rnd.masterList.Len(), "Master list size should be 3", t)
	assertEqualI(0, rnd.options.Len(), "Options should be empty at first", t)

	_ = rnd.Next()
	assertEqualI(3, rnd.masterList.Len(), "Master list size should still be 3", t)
	assertEqualI(2, rnd.options.Len(), "Options should include the remaining two strings", t)

	_ = rnd.Next()
	_ = rnd.Next()
	assertEqualI(0, rnd.options.Len(), "Options are empty after two more pulls", t)

	_ = rnd.Next()
	assertEqualI(2, rnd.options.Len(), "Options are refilled when needed", t)
}

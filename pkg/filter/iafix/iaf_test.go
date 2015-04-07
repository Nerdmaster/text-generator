package iafix // import "go.nerdbucket.com/text/pkg/filter/iafix"

import (
	"testing"
)

func assertEqualS(expected, actual string, message string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %#v, but got %#v - %s", expected, actual, message)
	}
}

func TestIAF(t *testing.T) {
	iaf := New()
	input := `I have a/an roc and a/an orc`
	assertEqualS("I have a roc and an orc", iaf.Filter(input), "Filter returns replaced text", t)
}

func TestSpanningNewlines(t *testing.T) {
	iaf := New()
	input := `Oh no, good sirs, a/an
	8-year-old is beating up Nerdmaster again!`
	assertEqualS(
		"Oh no, good sirs, an\n\t8-year-old is beating up Nerdmaster again!",
		iaf.Filter(input), "Filter properly spans lines", t,
	)
}

func TestCase(t *testing.T) {
	iaf := New()
	input := `A/an orange`
	assertEqualS("An orange", iaf.Filter(input), "Filter properly capitalizes", t)

	input = `I said, "A/AN orange"`
	assertEqualS(`I said, "AN orange"`, iaf.Filter(input), "Filter properly capitalizes AGAIN!", t)
}

package variation

import (
	"fmt"
	"math/rand"
	"testing"
)

func assertEqualS(expected, actual string, message string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %#v, but got %#v - %s", expected, actual, message)
	}
}

func TestSimple(t *testing.T) {
	v := New()
	input := `{{foo|bar|baz}}`
	outputs := []string{"baz", "bar", "foo", "foo", "foo", "baz"}

	var seed int64
	for index, s := range outputs {
		seed = int64(index+1)
		rand.Seed(seed)
		assertEqualS(s, v.Filter(input), fmt.Sprintf("Random text seed %d", seed), t)
	}
}

func TestNested(t *testing.T) {
	v := New()

	// WOW!  Using the commit message from BEFORE this code!  AMAZING!
	input := `...so that we can add a {{really|incredibly|s{{u|ooooo}}per|ultra}} awesome new filter!  It's {{gonna|going to}} be {{just swell|s{{|oo|ooo|oooo}} great}}!`

	outputs := []string{
		`...so that we can add a incredibly awesome new filter!  It's gonna be soooo great!`,
		`...so that we can add a super awesome new filter!  It's gonna be just swell!`,
		`...so that we can add a ultra awesome new filter!  It's going to be sooo great!`,
		`...so that we can add a incredibly awesome new filter!  It's going to be soo great!`,
		`...so that we can add a incredibly awesome new filter!  It's going to be just swell!`,
		`...so that we can add a ultra awesome new filter!  It's gonna be soooo great!`,
	}

	var seed int64
	for index, s := range outputs {
		seed = int64(index+1)
		rand.Seed(seed)
		assertEqualS(s, v.Filter(input), fmt.Sprintf("Random text seed %d", seed), t)
	}
}

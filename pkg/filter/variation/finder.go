package variation

import (
	"regexp"
	"strings"
)

// Format for variations - at least one pipe must be present, and text between
// pipes becomes the possible values.
var pattern = regexp.MustCompile(`{{(([^{}]*\|[^{}]*)+)}}`)

// The finder is responsible for finding variations and pulling out the pieces
// for the filter to use
type finder struct {
	source   *string
	options  []string
	fullText string
}

// Returns a new finder instance, holding a pointer to the source string so
// that as it changes, calls to Find continue to work
func makeFinder(text *string) *finder {
	return &finder{source: text}
}

func (f *finder) buildOptions() {
	m := pattern.FindStringSubmatch(*f.source)
	if m == nil {
		f.options = nil
		return
	}

	f.fullText = m[0]
	f.options = strings.Split(m[1], "|")
}

// Looks for the next occurrence of a variation.  If nothing is found, return
// value is false and no further actions should be taken.  If a match is found,
// its data is stored for the filter to use.
func (f *finder) Find() bool {
	f.buildOptions()
	return f.options != nil
}

// Returns the options last found
func (f *finder) Options() []string {
	return f.options
}

// Returns the full text of the last variation match found
func (f *finder) FullText() string {
	return f.fullText
}

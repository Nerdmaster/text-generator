package substitution // import "nerdbucket.com/go/text-generator/pkg/filter/substitution"

import "regexp"

// Format for substitutions
var pattern = regexp.MustCompile(`{{([^}]*)}}`)

// The finder is responsible for finding substitutions and pulling out the
// pieces for the filter to use
type finder struct {
	source   *string
	match    match
	fullText string
}

// Returns a new finder instance, holding a pointer to the source string so
// that as it changes, calls to Find continue to work
func makeFinder(text *string) *finder {
	return &finder{source: text}
}

func (f *finder) buildMatchData() {
	m := pattern.FindStringSubmatch(*f.source)
	if m == nil {
		f.match = emptyMatch
	} else {
		f.match = newMatch(m[1])
		f.fullText = m[0]
	}
}

// Looks for the next occurrence of a substitution.  If nothing is found,
// return value is false and no further actions should be taken.  If a match is
// found, its data is stored for the filter to use.
func (f *finder) Find() bool {
	f.buildMatchData()
	return f.match != emptyMatch
}

// Returns the last identifier found
func (f *finder) ID() string {
	return f.match.ID
}

// Returns the last variable name found
func (f *finder) VarName() string {
	return f.match.VarName
}

// Returns the full text of the last substitution match found
func (f *finder) FullText() string {
	return f.fullText
}

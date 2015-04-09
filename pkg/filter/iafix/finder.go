package iafix // import "go.nerdbucket.com/text/pkg/filter/iafix"

import "regexp"

// Format for substitutions
var pattern = regexp.MustCompile(`(?i)(a/an)(\s+)(\w+)`)

// The finder is responsible for finding indefinite article hacks ("a/an") and
// pulling out the pieces for the filter to use
type finder struct {
	source     *string
	fullText   string
	iaText     string
	whitespace string
	word       string
}

// Returns a new finder instance, holding a pointer to the source string so
// that as it changes, calls to Find continue to work
func makeFinder(text *string) *finder {
	return &finder{source: text}
}

func (f *finder) buildMatchData() {
	m := pattern.FindStringSubmatch(*f.source)
	if m == nil {
		f.word = ""
		return
	}

	f.fullText = m[0]
	f.iaText = m[1]
	f.whitespace = m[2]
	f.word = m[3]
}

// Looks for the next occurrence of a substitution.  If nothing is found,
// return value is false and no further actions should be taken.  If a match is
// found, its data is stored for the filter to use.
func (f *finder) Find() bool {
	f.buildMatchData()
	return f.word != ""
}

// Returns the "a/an" text as it was seen (including case)
func (f *finder) IndefiniteArticleText() string {
	return f.iaText
}

// Returns the last post-indefinite-article word found
func (f *finder) Word() string {
	return f.word
}

// Returns the last whitespace block found
func (f *finder) Whitespace() string {
	return f.whitespace
}

// Returns the full text of the last substitution match found
func (f *finder) FullText() string {
	return f.fullText
}

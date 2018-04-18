package variation // import "github.com/Nerdmaster/text-generator/pkg/filter/variation"

import (
	"github.com/Nerdmaster/text-generator/pkg/stringlist"
	"strings"
)

// Variation is a type of filter for picking one item out of a list.  Lists are
// surrounded by double-curly-braces and must have at least one pipe character
// (|) in them.  e.g., "{{foo|bar}}" will randomly choose either "foo" or "bar"
// in the text being examined.  The required pipe character means Variation can
// have a familiar syntax without conflicting with the substitution filter.
//
// Variation implements Filterable for use in templates.
type Variation struct{}

// New returns a new Variation with the null generator factory preset
func New() *Variation {
	return &Variation{}
}

// Filter finds all double-curly-brace tokens and replaces them with a value
// from the appropriate generator
func (s *Variation) Filter(text string) string {
	f := makeFinder(&text)

	for f.Find() {
		options := stringlist.FromSlice(f.Options())
		options.Shuffle()
		text = strings.Replace(text, f.FullText(), options.Pop(), 1)
	}

	return text
}

package template // import "nerdbucket.com/go/text-generator/lib/template"

import (
	"io/ioutil"
	"nerdbucket.com/go/text-generator/lib/filter"
)

type filterList []filter.Filterable

// A Template is a container for a text value and one or more filters
// (implementing filter.Filterable) the text should be run through.  The
// template itself doesn't have to be in any specific format, so long as it
// makes sense for whatever filters are attached to it.
//
// The order of a template's filters matter.  For instance:
//
// - a template is create with the string "Always add an extra article"
// - filter A replaces all occurrences of uppercase "A" with "@"
// - filter B uppercases all letters in a string
//
// If you add filter A, then filter B, template.Execute() will return:
//     "@LWAYS ADD AN EXTRA ARTICLE"
//
// If you add B first, then A, template.Execute() will instead return:
//     "@LW@YS @DD @N EXTR@ @RTICLE"
type Template struct {
	Text    string
	Filters filterList
}

// New returns an empty template with no text or filters
func New() *Template {
	return &Template{Filters: make(filterList, 0)}
}

// FromString sets up a new template using the given string as its source text
func FromString(text string) *Template {
	t := New()
	t.Text = text
	return t
}

// FromFile sets up and returns a new template by reading the given filename
// and converting its contents into a string
func FromFile(filename string) (*Template, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return FromString(string(fileBytes)), nil
}

// AddFilter puts a filter into the list of those run on template execution
func (t *Template) AddFilter(f filter.Filterable) {
	t.Filters = append(t.Filters, f)
}

// Execute runs through all filters in sequence
func (t *Template) Execute() string {
	out := t.Text
	for _, f := range t.Filters {
		out = f.Filter(out)
	}

	return out
}

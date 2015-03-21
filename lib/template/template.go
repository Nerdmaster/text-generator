package template // import "nerdbucket.com/go/text-generator/lib/template"

import (
	"fmt"
	"io/ioutil"
	"nerdbucket.com/go/text-generator/lib/filter"
)

type FilterList []filter.Filterable

type Template struct {
	Log     func(string)
	Text    string
	Filters FilterList
}

func dumblog(s string) {
	fmt.Printf(s)
}

func New() *Template {
	return &Template{Log: dumblog, Filters: make(FilterList, 0)}
}

func FromString(text string) *Template {
	t := New()
	t.Text = text
	return t
}

func FromFile(filename string) (*Template, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return FromString(string(fileBytes)), nil
}

func (t *Template) AddFilter(f filter.Filterable) {
	t.Filters = append(t.Filters, f)
}

// Reads the template and populate data
func (t *Template) Execute() string {
	out := t.Text
	for _, f := range t.Filters {
		out = f.Filter(out)
	}

	return out
}

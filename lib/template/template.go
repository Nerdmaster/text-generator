package template // import "nerdbucket.com/go/text-generator/lib/template"

import (
	"fmt"
	"io/ioutil"
	"nerdbucket.com/go/text-generator/lib/stringlist"
	"strings"
)

type Template struct {
	Log        func(string)
	Text       string
	Generators stringlist.GeneratorMap
}

func dumblog(s string) {
	fmt.Printf(s)
}

func New() *Template {
	return &Template{Log: dumblog, Generators: make(stringlist.GeneratorMap)}
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

// Pulls the next string for the requested generator, returning an error and ""
// if no generator exists
func (t *Template) GenerateString(name string) (string, error) {
	generator := t.Generators[name]
	if generator == nil {
		return "", fmt.Errorf("No generator named '%s' exists", name)
	}

	return generator.Next(), nil
}

// Reads the template and populate data
func (t *Template) Execute() string {
	sf := NewSubstitutionReplacer(t.Text)

	for sf.Find() {
		name, variable := splitSubstitutionIdentifier(sf.Identifier())
		value, err := t.GenerateString(name)
		if err != nil {
			t.Log(fmt.Sprintf("ERROR: %s", err))
		}
		t.AssignVariable(variable, value)
		sf.Replace(value)
	}

	return sf.Text()
}

// Stores a new SingleValueGenerator for the given variable name to return the
// specified value.  Nothing happens if variable is a blank string.
func (t *Template) AssignVariable(variable, value string) {
	if variable != "" {
		t.Generators[variable] = &SingleValueGenerator{Value: value}
	}
}

// Converts a substitution identifier into the generator name and possibly a
// variable name in which to store the generated data
func splitSubstitutionIdentifier(identifier string) (string, string) {
	data := strings.Split(identifier, "->")
	if len(data) < 2 {
		return data[0], ""
	}

	return data[0], data[1]
}

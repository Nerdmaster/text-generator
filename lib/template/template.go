package template // import "nerdbucket.com/go/text-generator/lib/template"

import (
	"fmt"
	"io/ioutil"
	"nerdbucket.com/go/text-generator/lib/stringlist"
	"regexp"
	"strings"
)

type Template struct {
	Log func(string)
	Text string
	Generators stringlist.GeneratorMap
}

var tvarRegex = regexp.MustCompile(`{{([^}]*)}}`)

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

// Reads the template and populate data
func (t *Template) Execute() string {
	out := t.Text
	for {
		foundStrings := tvarRegex.FindStringSubmatch(out)
		if foundStrings == nil {
			break
		}

		// Set up a variable to hold the replacement value
		replacementValue := ""

		// Store the full match in an alias for easier replacing later
		fullMatch := foundStrings[0]

		// Handle possible variable assignments
		data := strings.Split(foundStrings[1], "->")
		generatorName := data[0]
		variable := ""
		if len(data) == 2 {
			variable = data[1]
		}

		// See if the generator exists and warn if not
		generator := t.Generators[generatorName]
		if generator == nil {
			t.Log(fmt.Sprintf("ERROR: Generator '%s' needed but doesn't exist\n", generatorName))
		} else {
			replacementValue = generator.Next()
		}

		if variable != "" {
			t.Generators[variable] = &SingleValueGenerator{Value: replacementValue}
		}

		out = strings.Replace(out, fullMatch, replacementValue, 1)
	}

	return out
}

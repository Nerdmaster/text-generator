package substitution // import "nerdbucket.com/go/text-generator/lib/filter/substitution"

import (
	"nerdbucket.com/go/text-generator/lib/generator"
	"strings"
)

// The null generator is used when generator lookup fails, letting us return a
// blank string instead of otherwise failing.  This can be overridden to allow
// for other behaviors.
var MakeNullGenerator = func(id string) generator.Generator {
	return &generator.Static{Value: ""}
}

// Substitution is a type of filter built to convert special tokens wrapped in
// double-curly-braces into text via a Generator.  Implements the Filterable
// interface.
type Substitution struct {
	generators           generator.Map
	NullGeneratorFactory func(string) generator.Generator
}

func New() *Substitution {
	return &Substitution{generators: make(generator.Map), NullGeneratorFactory: MakeNullGenerator}
}

// Stores a new generator in the generators map, ignoring requests for setting
// a generator on ""
func (s *Substitution) SetGenerator(name string, g generator.Generator) {
	if name == "" {
		return
	}

	s.generators[name] = g
}

// Stores a static value for variables and other "generate once" situations.
func (s *Substitution) SetValue(name, value string) {
	s.SetGenerator(name, &generator.Static{Value: value})
}

// Returns the named generator or nullGenerator if no generator exists for the
// given id
func (s *Substitution) GetGenerator(id string) generator.Generator {
	g := s.generators[id]
	if g == nil {
		g = s.NullGeneratorFactory(id)
	}

	return g
}

func (s *Substitution) Filter(text string) string {
	f := makeFinder(&text)

	for f.Find() {
		replacement := s.GetGenerator(f.Id()).Next()
		s.SetValue(f.VarName(), replacement)
		text = strings.Replace(text, f.FullText(), replacement, 1)
	}

	return text
}

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
// interface for use in templates.
//
// The Substitution filter uses a finder that recognizes the "stabby" syntax
// for setting a static value.  The syntax must be precisely
// {{NAME->$VARIABLE}}.  i.e., if you have {{foo->$foo}}, the value will be
// pulled from a generator named "foo" via the GetGenerator function, and then
// it will assign a "$foo" as a static generator that always returns that same
// value.  This allows for a substitution with a certain level of context.  For
// instance:
//
//     {{boyname->$boy}}: Hello, my name is {{$boy}}
//     {{girlname->$girl}}: Hi, {{$boy}}, I'm {{$girl}}
//     {{boyname}}: You guys are lame, I'm gonna go to a different party
//
// This will result in all occurrences of {{$boy}} being the same "boyname" as
// the first one generated, while the last line will use a different "boyname"
// value (assuming the assigned generator has an extra value).
type Substitution struct {
	generators           generator.Map
	NullGeneratorFactory func(string) generator.Generator
}

// New returns a new Substitution with the null generator factory preset
func New() *Substitution {
	return &Substitution{generators: make(generator.Map), NullGeneratorFactory: MakeNullGenerator}
}

// SetGenerator stores a new generator in the generators map, ignoring requests for setting
// a generator on ""
func (s *Substitution) SetGenerator(name string, g generator.Generator) {
	if name == "" {
		return
	}

	s.generators[strings.ToLower(name)] = g
}

// SetValue stores a static value for variables and other "generate once" situations.
func (s *Substitution) SetValue(name, value string) {
	s.SetGenerator(name, &generator.Static{Value: value})
}

// GetGenerator returns the named generator or nullGenerator if no generator exists for the
// given id
func (s *Substitution) GetGenerator(id string) generator.Generator {
	g := s.generators[strings.ToLower(id)]
	if g == nil {
		g = s.NullGeneratorFactory(id)
	}

	return g
}

// fixCase returns a fixed version of s depending on id:
//
//     - If the id is all-title-cased, the return will be all-title-cased
//     - If the id's first rune is title-cased (uppercase), the return will have its first rune title-cased
//     - If none of the above are true, the return will be unmodified
func fixCase(id, s string) string {
	if strings.ToTitle(id) == id {
		return strings.ToTitle(s)
	}

	// Check the *second* rune for variables, which start with "$"
	r := id[:1]
	if r == "$" {
		r = id[1:2]
	}

	if strings.Title(r) == r {
		return strings.Title(s[:1]) + s[1:]
	}

	return s
}

// Filter finds all double-curly-brace tokens and replaces them with a value
// from the appropriate generator
func (s *Substitution) Filter(text string) string {
	f := makeFinder(&text)

	for f.Find() {
		replacement := s.GetGenerator(f.ID()).Next()
		replacement = fixCase(f.ID(), replacement)
		s.SetValue(f.VarName(), replacement)
		text = strings.Replace(text, f.FullText(), replacement, 1)
	}

	return text
}

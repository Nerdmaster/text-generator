package template // import "nerdbucket.com/go/text-generator/lib/template"

import (
	"regexp"
	"strings"
)

// Template substitution format is double-curly-braces around text
var templateSubstitutionRegex = regexp.MustCompile(`{{([^}]*)}}`)

// The SubstitutionReplacer handles finding substitutions in a string, and
// replacing those with arbitrary text
type SubstitutionReplacer struct {
	text             string
	identifier       string
	substitutionText string
}

func NewSubstitutionReplacer(text string) *SubstitutionReplacer {
	return &SubstitutionReplacer{text: text}
}

// Returns the underlying identifier (text between curly braces) from the most
// recent match found by calling Find()
func (sf *SubstitutionReplacer) Identifier() string {
	return sf.identifier
}

// Returns the text being searched / replaced
func (sf *SubstitutionReplacer) Text() string {
	return sf.text
}

// Looks for the next occurrence of a substitution.  If nothing is found,
// return value is false and no further actions should be taken.  If a match
// is found, its identifier is stored,
func (sf *SubstitutionReplacer) Find() bool {
	sf.substitutionText = ""
	sf.identifier = ""

	matches := templateSubstitutionRegex.FindStringSubmatch(sf.text)
	if matches == nil {
		return false
	}

	sf.substitutionText = matches[0]
	sf.identifier = matches[1]
	return true
}

func (sf *SubstitutionReplacer) Replace(replacement string) {
	sf.text = strings.Replace(sf.text, sf.substitutionText, replacement, 1)
}

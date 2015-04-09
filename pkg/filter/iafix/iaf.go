package iafix // import "go.nerdbucket.com/text/pkg/filter/iafix"

import (
	"fmt"
	"strings"
)

// IndefiniteArticleFix is a very simple filter solely for replacing occurrences
// of "a/an" with the proper indefinite article
type IndefiniteArticleFix struct{}

func New() *IndefiniteArticleFix {
	return &IndefiniteArticleFix{}
}

// This is an extremely basic fix that doesn't handle a huge number of edge
// cases, such as acronyms, words starting with a silent "H", possible UTF-8
// variances I know nothing about, etc.
func GetIndefiniteArticleFor(word string) string {
	start := strings.ToTitle(word[0:1])
	if strings.Contains("AEIOU8", start) {
		return "an"
	}
	return "a"
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

	r := id[:1]
	if strings.Title(r) == r {
		return strings.Title(s[:1]) + s[1:]
	}

	return s
}

// Filter finds all double-curly-brace tokens and replaces them with a value
// from the appropriate generator
func (iaf *IndefiniteArticleFix) Filter(text string) string {
	f := makeFinder(&text)

	for f.Find() {
		indefiniteArticle := GetIndefiniteArticleFor(f.Word())
		indefiniteArticle = fixCase(f.IndefiniteArticleText(), indefiniteArticle)
		replacement := fmt.Sprintf("%s%s%s", indefiniteArticle, f.Whitespace(), f.Word())
		text = strings.Replace(text, f.FullText(), replacement, -1)
	}

	return text
}

package substitution // import "nerdbucket.com/go/text-generator/lib/filter/substitution"

import "strings"

// Match represents the data for a variable-enabled matched structure
type match struct {
	ID      string
	VarName string
}

var emptyMatch = match{}

// Sets up a new match instance suitable for use when replacing text.
// Automatically detects and stores VarName if a variable is set.
func newMatch(id string) match {
	m := match{ID: id}
	m.splitVarID()
	return m
}

// Splits the ID on "->" to check for variable assignment
func (m *match) splitVarID() {
	data := strings.Split(m.ID, "->")
	if len(data) == 2 {
		m.ID = data[0]
		m.VarName = data[1]
	}
}

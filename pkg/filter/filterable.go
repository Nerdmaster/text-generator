package filter // import "github.com/Nerdmaster/text-generator/pkg/filter"

// Filterable types are used to convert text from one state into another
type Filterable interface {
	Filter(string) string
}

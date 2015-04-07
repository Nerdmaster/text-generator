package filter // import "go.nerdbucket.com/text/pkg/filter"

// Filterable types are used to convert text from one state into another
type Filterable interface {
	Filter(string) string
}

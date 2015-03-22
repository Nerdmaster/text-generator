package filter // import "nerdbucket.com/go/text-generator/lib/filter"

// Filterable types are used to convert text from one state into another
type Filterable interface {
	Filter(string) string
}

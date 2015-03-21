package filter // import "nerdbucket.com/go/text-generator/lib/filter"

type Filterable interface {
	Filter(string) string
}

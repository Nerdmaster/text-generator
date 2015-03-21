package template

// Simple stringlist.Generator implementation for handling variables
type SingleValueGenerator struct {
	Value string
}

func (s *SingleValueGenerator) Next() string {
	return s.Value
}

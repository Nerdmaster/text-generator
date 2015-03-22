package generator

// A Generator is any type that generates strings.  See the Random type for an
// example where each call to Next() returns a random string pulled from a
// master string list.
type Generator interface {
	Next() string
}

// Map is a simple string-to-generator list primarily for use in the
// substitution filter where we map from an identifier to a random string list
type Map map[string]Generator

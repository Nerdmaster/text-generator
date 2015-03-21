package generator

type Generator interface {
	Next() string
}

type Map map[string]Generator

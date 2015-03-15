package stringlist

type Generator interface {
	Next() string
}

type GeneratorMap map[string]Generator

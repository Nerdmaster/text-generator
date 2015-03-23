package substitution // import "nerdbucket.com/go/text-generator/pkg/filter/substitution"

import (
	"math/rand"
	"nerdbucket.com/go/text-generator/pkg/generator"
	"testing"
)

func assertEqualS(expected, actual string, message string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %#v, but got %#v - %s", expected, actual, message)
	}
}

func TestWithStaticValue(t *testing.T) {
	sub := New()
	input := `Testing {{one}}, {{two}}`
	sub.SetValue("one", "this is one")
	sub.SetValue("two", "this is two")

	assertEqualS("Testing this is one, this is two", sub.Filter(input), "Filter returns replaced text", t)
}

func TestWithStringlist(t *testing.T) {
	sub := New()
	input := `Testing {{one}}, {{two}}`
	listOne := generator.MakeRandom()
	listOne.Append("item 1.1")
	listOne.Append("item 1.2")
	listOne.Append("item 1.3")

	listTwo := generator.MakeRandom()
	listTwo.Append("item 2.1")
	listTwo.Append("item 2.2")
	listTwo.Append("item 2.3")

	sub.SetGenerator("one", listOne)
	sub.SetGenerator("two", listTwo)

	rand.Seed(1)
	assertEqualS("Testing item 1.3, item 2.1", sub.Filter(input), "Random text seed 1", t)
	rand.Seed(2)
	assertEqualS("Testing item 1.2, item 2.2", sub.Filter(input), "Random text seed 2", t)
	rand.Seed(3)
	assertEqualS("Testing item 1.1, item 2.3", sub.Filter(input), "Random text seed 3", t)
}

func TestWithVariables(t *testing.T) {
	sub := New()
	input := `Testing {{one}}, {{two->$two}}, {{$two}}, {{$two}}`
	listOne := generator.MakeRandom()
	listOne.Append("item 1.1")
	listOne.Append("item 1.2")
	listOne.Append("item 1.3")

	listTwo := generator.MakeRandom()
	listTwo.Append("item 2.1")
	listTwo.Append("item 2.2")
	listTwo.Append("item 2.3")

	sub.SetGenerator("one", listOne)
	sub.SetGenerator("two", listTwo)

	rand.Seed(1)
	assertEqualS("Testing item 1.3, item 2.1, item 2.1, item 2.1", sub.Filter(input), "Random text vars seed 1", t)
	rand.Seed(2)
	assertEqualS("Testing item 1.2, item 2.2, item 2.2, item 2.2", sub.Filter(input), "Random text vars seed 2", t)
	rand.Seed(3)
	assertEqualS("Testing item 1.1, item 2.3, item 2.3, item 2.3", sub.Filter(input), "Random text vars seed 3", t)
}

func TestWithNesting(t *testing.T) {
	sub := New()
	input := `Testing {{one}}, {{two}}`
	sub.SetValue("one", "this is one")
	sub.SetValue("two", "this is two or {{three}}")
	sub.SetValue("three", "{{number3}}")
	sub.SetValue("number3", "3")
	assertEqualS("Testing this is one, this is two or 3", sub.Filter(input), "Filter replaces nested text", t)
}

func TestCase(t *testing.T) {
	sub := New()
	sub.SetValue("TeSt", "this is my test")
	assertEqualS("this is my test", sub.Filter("{{test}}"), "Filter stored 'TeSt' without case", t)
	assertEqualS("This is my test", sub.Filter("{{Test}}"), "Uppercased first rune", t)
	assertEqualS("THIS IS MY TEST", sub.Filter("{{TEST}}"), "Uppercased whole string", t)
}

func TestCaseForVariables(t *testing.T) {
	sub := New()
	input := `{{test->$BlAh}} {{$blah}} {{$Blah}} {{$BLAH}}`
	sub.SetValue("test", "test")
	assertEqualS("test test Test TEST", sub.Filter(input), "Variable cases", t)
}

package generator

import "nerdbucket.com/go/text-generator/lib/stringlist"

// The Random generator produces a random sequence of a given stringlist, using
// each string in the list exactly once before any single string is repeated
type Random struct {
	masterList *stringlist.List
	options    *stringlist.List
}

// MakeRandom returns a new empty Random generator
func MakeRandom() *Random {
	return &Random{
		masterList: stringlist.New(0),
		options:    stringlist.New(0),
	}
}

// Append adds a single string to the generator's master list
func (r *Random) Append(str string) {
	r.masterList.Append(str)
}

// Next retrieves the next string from the list of possible options, re-cloning
// options from the master list once all options have been used once
func (r *Random) Next() string {
	if r.options.Len() < 1 {
		r.options = r.masterList.Clone()
		r.options.Shuffle()
	}

	return r.options.Pop()
}

// IsEmpty returns whether or not the master list has any strings
func (r *Random) IsEmpty() bool {
	return r.masterList.Len() == 0
}

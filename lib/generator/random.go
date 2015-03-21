package generator

import "nerdbucket.com/go/text-generator/lib/stringlist"

type Random struct {
	masterList *stringlist.List
	options    *stringlist.List
}

func MakeRandom() *Random {
	return &Random{
		masterList: stringlist.New(0),
		options:    stringlist.New(0),
	}
}

func (r *Random) Append(str string) {
	r.masterList.Append(str)
}

func (r *Random) Next() string {
	// Clone and shuffle the master list if we have no strings
	if r.options.Len() < 1 {
		r.options = r.masterList.Clone()
		r.options.Shuffle()
	}

	return r.options.Pop()
}

func (r *Random) IsEmpty() bool {
	return r.masterList.Len() == 0
}

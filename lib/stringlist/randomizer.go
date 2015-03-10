package stringlist

type Randomizer struct {
	masterList *List
	options    *List
}

func MakeRandomizer() *Randomizer {
	return &Randomizer{
		masterList: New(0),
		options:    New(0),
	}
}

func (r *Randomizer) Append(str string) {
	r.masterList.Append(str)
}

func (r *Randomizer) Next() string {
	// Clone and shuffle the master list if we have no strings
	if r.options.Len() < 1 {
		r.options = r.masterList.Clone()
		r.options.Shuffle()
	}

	return r.options.Pop()
}

func (r *Randomizer) IsEmpty() bool {
	return r.masterList.Len() == 0
}

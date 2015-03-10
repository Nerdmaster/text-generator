package stringlist // import "nerdbucket.com/go/text-generator/lib/stringlist"

import "math/rand"

type List struct {
	data []string
}

func New(size int) *List {
	return &List{data: make([]string, size)}
}

func (list *List) Append(item string) {
	list.data = append(list.data, item)
}

func (list *List) Pop() string {
	size := list.Len() - 1

	if size == -1 {
		panic("Trying to pop from an empty List")
	}

	str := list.data[size]
	list.data = list.data[:size]

	return str
}

func (list *List) Len() int {
	return len(list.data)
}

func (list *List) Shuffle() {
	for i := range list.data {
		j := rand.Intn(i + 1)
		list.data[i], list.data[j] = list.data[j], list.data[i]
	}
}

func (list *List) Clone() *List {
	newlist := New(list.Len())
	copy(newlist.data, list.data)

	return newlist
}

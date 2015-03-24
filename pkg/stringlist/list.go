package stringlist // import "go.nerdbucket.com/text/pkg/stringlist"

import "math/rand"

// List is a very simple container for a slice of strings
type List struct {
	data []string
}

// New returns an empty list
func New(size int) *List {
	return &List{data: make([]string, size)}
}

// Append adds a string to the end of the list
func (list *List) Append(item string) {
	list.data = append(list.data, item)
}

// Pop removes the last item from the list and returns it
func (list *List) Pop() string {
	size := list.Len() - 1

	if size == -1 {
		panic("Trying to pop from an empty List")
	}

	str := list.data[size]
	list.data = list.data[:size]

	return str
}

// Len returns the number of items in the list
func (list *List) Len() int {
	return len(list.data)
}

// Shuffle randomly swaps each item in the list with another
func (list *List) Shuffle() {
	for i := range list.data {
		j := rand.Intn(i + 1)
		list.data[i], list.data[j] = list.data[j], list.data[i]
	}
}

// Clone creates a copy of the list as another list pointer
func (list *List) Clone() *List {
	newlist := New(list.Len())
	copy(newlist.data, list.data)

	return newlist
}

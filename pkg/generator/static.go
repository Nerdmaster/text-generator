package generator

// Static is a simple Generator implementation for handling static values
type Static struct {
	Value string
}

// Next returns the static value Value every time it's called
func (v *Static) Next() string {
	return v.Value
}

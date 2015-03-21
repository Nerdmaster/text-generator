package generator

// Static is a simple Generator implementation for handling static values
type Static struct {
	Value string
}

func (v *Static) Next() string {
	return v.Value
}

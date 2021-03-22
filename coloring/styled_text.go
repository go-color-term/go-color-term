package coloring

// StyledText represents a string decorated with style attributes.
//
// It can be obtained from a `StyleBuilder` or a `SentenceBuilder`.
type StyledText struct {
	s, unformatted string
}

// String returns the string represented by this `StyledText` with
// the corresponding style attributes applied during it's creation.
func (st *StyledText) String() string {
	return st.s
}

// Unformatted returns the original string without any style attributes.
func (st *StyledText) Unformatted() string {
	return st.unformatted
}

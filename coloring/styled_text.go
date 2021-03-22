package coloring

type StyledText struct {
	s, unformatted string
}

func (st *StyledText) String() string {
	return st.s
}

func (st *StyledText) Unformatted() string {
	return st.unformatted
}

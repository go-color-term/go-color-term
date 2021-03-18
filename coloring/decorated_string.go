package coloring

type DecoratedText struct {
	s, seq string
}

func (dt *DecoratedText) String() string {
	return applyTo(dt.seq, dt.s)
}

func (dt *DecoratedText) Unformatted() string {
	return dt.s
}

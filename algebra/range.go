package algebra

type Range struct {
	Start MnFloat
	End MnFloat
}

// Represents a range that covers infinity
var InfiniteRange = Range{
	Start: NegativeInfinity,
	End:   PositiveInfinity,
}

var EmptyRange = Range{MnZero, MnZero}

// Represents a range that is invalid
var InvalidRange = Range{
	Start: PositiveInfinity,
	End:   NegativeInfinity,
}

var FrontRange = Range{
	Start: Epsilon,
	End: PositiveInfinity,
}

var BackRange = Range{
	Start: NegativeInfinity,
	End: -Epsilon,
}

func (r Range) Contains(d MnFloat) bool {
	return d > r.Start && d < r.End
}

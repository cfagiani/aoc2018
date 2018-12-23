package datastructure

type IntPair struct {
	A int
	B int
}

func (a IntPair) Equals(b IntPair) bool {
	return a.A == b.A && a.B == b.B
}

//Compares 2 pairs. If aFirst is true, then the comparison will look at the A value first and only compare B values
//if A's are equal. If false, it will do the opposite
func (a IntPair) Less(b IntPair, aFirst bool) bool {
	if aFirst {
		if a.A == b.A {
			return a.B < b.B
		}
		return a.A < b.A
	}
	if a.B == b.B {
		return a.A < b.A
	}
	return a.B < b.B
}

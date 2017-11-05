package binoislist

// BinoisPtr is a basic type in order to be able to run unit tests.
type BinoisPtr int

func (b BinoisPtr) Equals(other BinoisPtr) bool {
	return b == other
}

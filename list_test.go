package binoislist

import "testing"

func TestSize(t *testing.T) {
	var list BinoisList

	if s := list.Len(); s != 0 {
		t.Errorf("Empty list wrong size, expected 0, got %d", s)
	}

	newList := list.Append(12)

	if s := newList.Len(); s != 1 {
		t.Errorf("Single item list wrong size, expected 1, got %d", s)
	}
}

func TestNewBinoisList(t *testing.T) {
	emptyList := NewBinoisList()

	if s := emptyList.Len(); s != 0 {
		t.Errorf("Empty list wrong size, expected 0, got %d", s)
	}

	list1 := NewBinoisList(1)
	if size := list1.Len(); size != 1 {
		t.Errorf("list1 wrong size, expected 1, got %d", size)
	}

	list2 := NewBinoisList(1, 2)
	if size := list2.Len(); size != 2 {
		t.Errorf("list2 wrong size, expected 2, got %d", size)
	}
}

func TestAppend(t *testing.T) {
	var list BinoisList

	list1 := list.Append(1)
	list2 := list1.Append(2)
	list3 := list2.Append(3)

	if size := list1.Len(); size != 1 {
		t.Errorf("list1 wrong size, expected 1, got %d", size)
	}
	if size := list2.Len(); size != 2 {
		t.Errorf("list2 wrong size, expected 2, got %d", size)
	}
	if size := list3.Len(); size != 3 {
		t.Errorf("list3 wrong size, expected 3, got %d", size)
	}

	list4 := list.Append(1, 2, 3, 4)

	if size := list4.Len(); size != 4 {
		t.Errorf("list4 wrong size, expected 4, got %d", size)
	}

	// Make sure previously created lists are still the same size.
	if size := list1.Len(); size != 1 {
		t.Errorf("list1 wrong size, expected 1, got %d", size)
	}
	if size := list2.Len(); size != 2 {
		t.Errorf("list2 wrong size, expected 2, got %d", size)
	}
	if size := list3.Len(); size != 3 {
		t.Errorf("list3 wrong size, expected 3, got %d", size)
	}
}

func TestJoinCat(t *testing.T) {
	aList := NewBinoisList(1, 2, 3)
	bList := NewBinoisList(4, 5, 6)

	cList := aList.Cat(bList)

	if size := cList.Len(); size != 6 {
		t.Errorf("expected joined list to be size 6, got %d", size)
	}

	aString := aList.Join(":")

	if aString != "1:2:3" {
		t.Errorf("expected \"1:2:3\" but got %s", aString)
	}

	bString := bList.Join("")

	if bString != "456" {
		t.Errorf("expected \"456\" but got %s", bString)
	}

	cString := cList.Join(", ")

	if cString != "1, 2, 3, 4, 5, 6" {
		t.Errorf("expected \"1, 2, 3, 4, 5, 6\" but got %s", cString)
	}

	dString := cList.Join()

	if dString != "1, 2, 3, 4, 5, 6" {
		t.Errorf("default separator should be \", \" yielding \"1, 2, 3, 4, 5, 6\" but got %s", dString)
	}
}
func TestGet(t *testing.T) {
	x := NewBinoisList(1, 2, 3)

	if n := x.Get(0); n != 1 {
		t.Errorf("expected 1, got %d", n)
	}
	if n := x.Get(1); n != 2 {
		t.Errorf("expected 2, got %d", n)
	}
	if n := x.Get(2); n != 3 {
		t.Errorf("expected 3, got %d", n)
	}
}

func TestGetOrUnknown(t *testing.T) {
	x := NewBinoisList(1)

	n, err := x.GetOrUnknown(0)
	if n != 1 {
		t.Errorf("expected 1, got %d", n)
	}
	if err != nil {
		t.Errorf("exepected no error, got %s", err)
	}

	n, err = x.GetOrUnknown(1)
	if n != UnknownBinois {
		t.Errorf("expected %d, got %d", UnknownBinois, n)
	}
	if err == nil || err.Error() != "out of range" {
		t.Errorf("exepected \"out of range\" error, got %s", err)
	}
}

func TestMap(t *testing.T) {
	x := NewBinoisList(1, 2, 3)

	y := x.Map(func(g BinoisPtr) BinoisPtr {
		return g + 100
	})

	if s := y.Len(); s != 3 {
		t.Errorf("expected size 3, got %d", s)
	}
	if n := y.Get(0); n != 101 {
		t.Errorf("expected 101, got %d", n)
	}
	if n := y.Get(1); n != 102 {
		t.Errorf("expected 102, got %d", n)
	}
	if n := y.Get(2); n != 103 {
		t.Errorf("expected 103, got %d", n)
	}

	var bigList BinoisList
	for i := 0; i < 10000; i++ {
		bigList = bigList.Append(BinoisPtr(i))
	}

	newBigList := bigList.Map(func(g BinoisPtr) BinoisPtr {
		return g + 100
	})

	if n := newBigList.Get(9999); n != 9999+100 {
		t.Errorf("expected 9999+100, got %d", n)
	}

}

func TestFilter(t *testing.T) {
	x := NewBinoisList(1, 2, 3)

	y := x.Filter(func(g BinoisPtr) bool {
		return g == 2
	})

	if y.Len() != 1 {
		t.Errorf("expected size 1, got %d", y.Len())
	}

	if y.Len() == 1 && y.Get(0) != 2 {
		t.Errorf("expected element to have value 2, got %d", y.Get(0))
	}
}

func TestFirst(t *testing.T) {
	x := NewBinoisList(1, 2, 3)

	y, found := x.First(func(g BinoisPtr) bool {
		return g == 2
	})

	if !found {
		t.Errorf("expected to find 2, but didn't")
	}

	if found && y != 2 {
		t.Errorf("expected to get value 2, got %d", y)
	}

	y, found = x.First(func(g BinoisPtr) bool {
		return g == 99
	})

	if found {
		t.Errorf("expected to not find 99, but got found")
	}

	if !found && y != UnknownBinois {
		t.Errorf("expected to unknown value %d, got %d", UnknownBinois, y)
	}
}

func TestCopy(t *testing.T) {
	a := NewBinoisList(1, 2, 3)

	b := a.Copy()

	if !a.Equals(b) {
		t.Errorf("expected copy of a list to be equal, but failed")
	}
}

func TestSort(t *testing.T) {
	a := NewBinoisList(1, 2, 3, 3, 2, 1)

	y := a.Copy().Sort(func(a, b BinoisPtr) bool {
		return a < b
	})

	sorted := NewBinoisList(1, 1, 2, 2, 3, 3)

	if !y.Equals(sorted) {
		t.Errorf("expected sort of [%s] to return [%s], but got [%s]", a.Join(), sorted.Join(), y.Join())
	}
}

func TestEquals(t *testing.T) {
	a := NewBinoisList(1, 2, 3)
	b := NewBinoisList(1, 2, 3)
	c := NewBinoisList(1, 3, 2)
	d := NewBinoisList(1, 2)

	y := a.Equals(b)

	if !y {
		t.Errorf("expected lists a and b to be equal, but got not equal")
	}

	y = a.Equals(c)

	if y {
		t.Errorf("expected lists a and c to not be equal, but got equal")
	}

	y = a.Equals(d)

	if y {
		t.Errorf("expected lists a and d to not be equal, but got equal")
	}

}

func TestIndex(t *testing.T) {
	a := NewBinoisList(1, 2, 3, 3, 2, 1)

	y, found := a.Index(3)

	if !found {
		t.Errorf("expected Index() to find 3, but didn't")
	}

	if found && y != 2 {
		t.Errorf("expected Index() to return 2, but got %d", y)
	}

	y, found = a.Index(1)

	if !found {
		t.Errorf("expected Index() to find 1, but didn't")
	}

	if found && y != 0 {
		t.Errorf("expected Index() to return 0, but got %d", y)
	}

	y, found = a.Index(99)

	if found {
		t.Errorf("expected Index() to not find 99, but got found")
	}

	if !found && y != -1 {
		t.Errorf("expected Index() to return %d, but got %d", -1, y)
	}
}

func TestLastIndex(t *testing.T) {
	a := NewBinoisList(1, 2, 3, 3, 2, 1)

	y, found := a.LastIndex(3)

	if !found {
		t.Errorf("expected LastIndex() to find 3, but didn't")
	}

	if found && y != 3 {
		t.Errorf("expected LastIndex() to return 3, but got %d", y)
	}

	y, found = a.LastIndex(1)

	if !found {
		t.Errorf("expected LastIndex() to find 1, but didn't")
	}

	if found && y != 5 {
		t.Errorf("expected LastIndex() to return 5, but got %d", y)
	}

	y, found = a.LastIndex(99)

	if found {
		t.Errorf("expected LastIndex() to not find 99, but got found")
	}

	if !found && y != -1 {
		t.Errorf("expected LastIndex() to return %d, but got %d", -1, y)
	}
}

func TestDelete(t *testing.T) {
	a := NewBinoisList(1, 2, 3)

	y := a.Delete(1)
	expected := NewBinoisList(1, 3)

	if !y.Equals(expected) {
		t.Errorf("expected Delete(1) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Delete(0)
	expected = NewBinoisList(2, 3)

	if !y.Equals(expected) {
		t.Errorf("expected Delete(0) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Delete(99)
	expected = NewBinoisList(1, 2, 3)

	if !y.Equals(expected) {
		t.Errorf("expected Delete(99) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Delete(-1)
	expected = NewBinoisList(1, 2)

	if !y.Equals(expected) {
		t.Errorf("expected Delete(-1) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Delete(-2)
	expected = NewBinoisList(1, 3)

	if !y.Equals(expected) {
		t.Errorf("expected Delete(-2) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Delete(-3)
	expected = NewBinoisList(2, 3)

	if !y.Equals(expected) {
		t.Errorf("expected Delete(-3) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Delete(-4)
	expected = NewBinoisList(1, 2, 3)

	if !y.Equals(expected) {
		t.Errorf("expected Delete(-4) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}
}

func TestPart(t *testing.T) {
	a := NewBinoisList(1, 2, 3, 4, 5)

	y := a.Part(0, 6)
	expected := a.Copy()

	if !y.Equals(expected) {
		t.Errorf("expected Part(0,6) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(0, 0)
	expected = NewBinoisList()

	if !y.Equals(expected) {
		t.Errorf("expected Part(0,0) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(0, 1)
	expected = NewBinoisList(1)

	if !y.Equals(expected) {
		t.Errorf("expected Part(0,1) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(2, 4)
	expected = NewBinoisList(3, 4)

	if !y.Equals(expected) {
		t.Errorf("expected Part(2,4) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(2, -1)
	expected = NewBinoisList(3, 4)

	if !y.Equals(expected) {
		t.Errorf("expected Part(2,-1) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(2, -2)
	expected = NewBinoisList(3)

	if !y.Equals(expected) {
		t.Errorf("expected Part(2,-2) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(2, -3)
	expected = NewBinoisList()

	if !y.Equals(expected) {
		t.Errorf("expected Part(2,-3) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(2, -4)
	expected = NewBinoisList()

	if !y.Equals(expected) {
		t.Errorf("expected Part(2,-4) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(-5, -4)
	expected = NewBinoisList(1)

	if !y.Equals(expected) {
		t.Errorf("expected Part(-5,-4) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(-5, -3)
	expected = NewBinoisList(1, 2)

	if !y.Equals(expected) {
		t.Errorf("expected Part(-5,-3) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}

	y = a.Part(-9, -3)
	expected = NewBinoisList(1, 2)

	if !y.Equals(expected) {
		t.Errorf("expected Part(-5,-3) to yield [%s], but got [%s]", expected.Join(), y.Join())
	}
}

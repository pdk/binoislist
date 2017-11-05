package binoislist

import (
	"fmt"
	"sort"
	"strings"
)

// BinoisList is a slice of BinoisPtr.
type BinoisList []BinoisPtr

const UnknownBinois BinoisPtr = -1

// NewBinoisList creates a new BinoisList with some initial items.
func NewBinoisList(items ...BinoisPtr) BinoisList {
	return items
}

// Append adds one or more items to the list.
// Don't forget to assign the result to your variable.
//   myList = myList.Append(x, y)
func (list BinoisList) Append(items ...BinoisPtr) BinoisList {
	return append(list, items...)
}

// Cat concatenates multiple lists into a single list.
// Don't forget to assign the result to your variable.
//   myList = myList.Cat(l1, l2)
func (list BinoisList) Cat(inputLists ...BinoisList) BinoisList {
	for _, oneList := range inputLists {
		list = append(list, oneList...)
	}
	return list
}

// Join converts all the items in a BinoisList to strings and joins them with a
// separator. Default separator is ", ".
func (list BinoisList) Join(seps ...string) string {
	sep := ", "
	if len(seps) > 0 {
		sep = seps[0]
	}

	var collector []string
	for _, item := range list {
		// %#v will invoke String() if exists on the item.
		collector = append(collector, fmt.Sprintf("%#v", item))
	}
	return strings.Join(collector, sep)
}

// Len returns the number of elements in the list
func (list BinoisList) Len() int {
	return len(list)
}

// Get returns the item at position i.
func (list BinoisList) Get(i int) BinoisPtr {
	return list[i]
}

// GetOrUnknown returns the element as position i, or unknown if out of range.
func (list BinoisList) GetOrUnknown(i int) (BinoisPtr, error) {
	if i < 0 || i >= len(list) {
		return UnknownBinois, fmt.Errorf("out of range")
	}
	return list[i], nil
}

// BinoisMapperFunc processes a BinoisPtr and returns a BinoisPtr.
type BinoisMapperFunc func(g BinoisPtr) BinoisPtr

func (list BinoisList) Map(f BinoisMapperFunc) BinoisList {
	var newList BinoisList
	for _, item := range list {
		newList = append(newList, f(item))
	}
	return newList
}

// BinoisTesterFunc is a function that returns true/false given a BinoisPtr. Used
// for finding or filtering with a BinoisList.
type BinoisTesterFunc func(g BinoisPtr) bool

// Filter returns a new BinoisList of items where f() is true.
func (list BinoisList) Filter(f BinoisTesterFunc) BinoisList {
	var newList BinoisList
	for _, item := range list {
		if f(item) {
			newList = append(newList, item)
		}
	}
	return newList
}

// First returns the first BinoisPtr where f() is true.
func (list BinoisList) First(f BinoisTesterFunc) (item BinoisPtr, found bool) {
	for _, item := range list {
		if f(item) {
			return item, true
		}
	}

	return UnknownBinois, false
}

// BinoisLessThanFunc is a function that compares two items of type BinoisPtr.
type BinoisLessThanFunc func(a, b BinoisPtr) bool

// Sort sorts IN PLACE a BinoisList with the input comparer function.
func (list BinoisList) Sort(f BinoisLessThanFunc) BinoisList {
	sort.Slice(list, func(i, j int) bool {
		return f(list[i], list[j])
	})

	return list
}

// Copy returns a BinoisList containing the contents of the list.
func (list BinoisList) Copy() BinoisList {
	newList := make(BinoisList, len(list))
	copy(newList, list)
	return newList
}

// BinoisEqualizer provides an Equals method for items of type BinoisPtr.
type BinoisEqualizer interface {
	Equals(BinoisPtr) bool
}

// Equalizer handles type conversion. This is a workaround to allow users
// of this package to partially implement requirements.
func (thing BinoisPtr) Equalizer() BinoisEqualizer {
	var x interface{} = thing
	return x.(BinoisEqualizer)
}

// affirmBinoisEqualsImplemented exists to output a more meaningful error when
// the Equals method has not been implemented. We could have just done type
// assertions w/o checking and had the same net result (panic), but having a
// better explanation is better.
func affirmBinoisEqualsImplemented(methodName string) {
	var t BinoisPtr
	var x interface{} = t
	_, ok := x.(BinoisEqualizer)
	if !ok {
		panic("implement method BinoisPtr.Equals(BinoisPtr) in order to use " + methodName)
	}
}

// Index returns the position  of item in the list, and if found.
// Method Equals(BinoisPtr) must be implemented.
func (list BinoisList) Index(item BinoisPtr) (int, bool) {
	affirmBinoisEqualsImplemented("BinoisList.Index(item)")

	eqItem := item.Equalizer()

	for pos, item := range list {
		if eqItem.Equals(item) {
			return pos, true
		}
	}

	return -1, false
}

// LastIndex returns the position  of item in the list, and if found.
// Method Equals(BinoisPtr) must be implemented.
func (list BinoisList) LastIndex(item BinoisPtr) (int, bool) {
	affirmBinoisEqualsImplemented("BinoisList.LastIndex(item)")

	eqItem := item.Equalizer()

	for pos := len(list) - 1; pos >= 0; pos-- {
		if eqItem.Equals(list[pos]) {
			return pos, true
		}
	}

	return -1, false
}

// Equals checks if two BinoisList instances are equal by comparing all their
// elements with the BinoisPtr Equals method.
func (list BinoisList) Equals(other BinoisList) bool {
	affirmBinoisEqualsImplemented("BinoisList.Equals()")

	if list.Len() != other.Len() {
		return false
	}

	for i := range list {
		eqItem := list[i].Equalizer()
		if !eqItem.Equals(other[i]) {
			return false
		}
	}

	return true
}

// Delete returns a new list with one element omitted.
func (list BinoisList) Delete(index int) BinoisList {
	if len(list)+index < 0 {
		return list.Copy()
	}
	index = list.cleanIndex(index)
	return list.DeletePart(index, index+1)
}

// cleanIndex is a helper method to ensure offsets are within range.
func (list BinoisList) cleanIndex(i int) int {
	if i > len(list) {
		return len(list)
	}
	if i < 0 {
		i = len(list) + i
		if i < 0 {
			return 0
		}
		return i
	}
	return i
}

// DeletePart returns a new BinoisList with a portion omitted.
func (list BinoisList) DeletePart(first, last int) BinoisList {
	first = list.cleanIndex(first)
	last = list.cleanIndex(last)

	if first >= last {
		return list.Copy()
	}

	var newList BinoisList
	newList = append(newList, list[0:first]...)
	newList = append(newList, list[last:]...)

	return newList
}

// Part returns a portion of a BinoisList as a new list.
func (list BinoisList) Part(first, last int) BinoisList {
	first = list.cleanIndex(first)
	last = list.cleanIndex(last)

	if first >= last {
		return NewBinoisList()
	}

	return list[first:last]
}

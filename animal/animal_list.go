// Code generated with make-binois-list.sh DO NOT EDIT.
// http://github.com/pdk/binoislist

package animal

import (
	"fmt"
	"sort"
	"strings"
)

// AnimalsList is a slice of Animal.
type AnimalsList []Animal


// NewAnimalsList creates a new AnimalsList with some initial items.
func NewAnimalsList(items ...Animal) AnimalsList {
	return items
}

// Append adds one or more items to the list.
// Don't forget to assign the result to your variable.
//   myList = myList.Append(x, y)
func (list AnimalsList) Append(items ...Animal) AnimalsList {
	return append(list, items...)
}

// Cat concatenates multiple lists into a single list.
// Don't forget to assign the result to your variable.
//   myList = myList.Cat(l1, l2)
func (list AnimalsList) Cat(inputLists ...AnimalsList) AnimalsList {
	for _, oneList := range inputLists {
		list = append(list, oneList...)
	}
	return list
}

// Join converts all the items in a AnimalsList to strings and joins them with a
// separator. Default separator is ", ".
func (list AnimalsList) Join(seps ...string) string {
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
func (list AnimalsList) Len() int {
	return len(list)
}

// Get returns the item at position i.
func (list AnimalsList) Get(i int) Animal {
	return list[i]
}

// GetOrUnknown returns the element as position i, or unknown if out of range.
func (list AnimalsList) GetOrUnknown(i int) (Animal, error) {
	if i < 0 || i >= len(list) {
		return Animal{}, fmt.Errorf("out of range")
	}
	return list[i], nil
}

// AnimalMapperFunc processes a Animal and returns a Animal.
type AnimalMapperFunc func(g Animal) Animal

func (list AnimalsList) Map(f AnimalMapperFunc) AnimalsList {
	var newList AnimalsList
	for _, item := range list {
		newList = append(newList, f(item))
	}
	return newList
}

// AnimalTesterFunc is a function that returns true/false given a Animal. Used
// for finding or filtering with a AnimalsList.
type AnimalTesterFunc func(g Animal) bool

// Filter returns a new AnimalsList of items where f() is true.
func (list AnimalsList) Filter(f AnimalTesterFunc) AnimalsList {
	var newList AnimalsList
	for _, item := range list {
		if f(item) {
			newList = append(newList, item)
		}
	}
	return newList
}

// First returns the first Animal where f() is true.
func (list AnimalsList) First(f AnimalTesterFunc) (item Animal, found bool) {
	for _, item := range list {
		if f(item) {
			return item, true
		}
	}

	return Animal{}, false
}

// AnimalLessThanFunc is a function that compares two items of type Animal.
type AnimalLessThanFunc func(a, b Animal) bool

// Sort sorts IN PLACE a AnimalsList with the input comparer function.
func (list AnimalsList) Sort(f AnimalLessThanFunc) AnimalsList {
	sort.Slice(list, func(i, j int) bool {
		return f(list[i], list[j])
	})

	return list
}

// Copy returns a AnimalsList containing the contents of the list.
func (list AnimalsList) Copy() AnimalsList {
	newList := make(AnimalsList, len(list))
	copy(newList, list)
	return newList
}

// AnimalEqualizer provides an Equals method for items of type Animal.
type AnimalEqualizer interface {
	Equals(Animal) bool
}

// Equalizer handles type conversion. This is a workaround to allow users
// of this package to partially implement requirements.
func (thing Animal) Equalizer() AnimalEqualizer {
	var x interface{} = thing
	return x.(AnimalEqualizer)
}

// affirmAnimalEqualsImplemented exists to output a more meaningful error when
// the Equals method has not been implemented. We could have just done type
// assertions w/o checking and had the same net result (panic), but having a
// better explanation is better.
func affirmAnimalEqualsImplemented(methodName string) {
	var t Animal
	var x interface{} = t
	_, ok := x.(AnimalEqualizer)
	if !ok {
		panic("implement method Animal.Equals(Animal) in order to use " + methodName)
	}
}

// Index returns the position  of item in the list, and if found.
// Method Equals(Animal) must be implemented.
func (list AnimalsList) Index(item Animal) (int, bool) {
	affirmAnimalEqualsImplemented("AnimalsList.Index(item)")

	eqItem := item.Equalizer()

	for pos, item := range list {
		if eqItem.Equals(item) {
			return pos, true
		}
	}

	return -1, false
}

// LastIndex returns the position  of item in the list, and if found.
// Method Equals(Animal) must be implemented.
func (list AnimalsList) LastIndex(item Animal) (int, bool) {
	affirmAnimalEqualsImplemented("AnimalsList.LastIndex(item)")

	eqItem := item.Equalizer()

	for pos := len(list) - 1; pos >= 0; pos-- {
		if eqItem.Equals(list[pos]) {
			return pos, true
		}
	}

	return -1, false
}

// Equals checks if two AnimalsList instances are equal by comparing all their
// elements with the Animal Equals method.
func (list AnimalsList) Equals(other AnimalsList) bool {
	affirmAnimalEqualsImplemented("AnimalsList.Equals()")

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
func (list AnimalsList) Delete(index int) AnimalsList {
	if len(list)+index < 0 {
		return list.Copy()
	}
	index = list.cleanIndex(index)
	return list.DeletePart(index, index+1)
}

// cleanIndex is a helper method to ensure offsets are within range.
func (list AnimalsList) cleanIndex(i int) int {
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

// DeletePart returns a new AnimalsList with a portion omitted.
func (list AnimalsList) DeletePart(first, last int) AnimalsList {
	first = list.cleanIndex(first)
	last = list.cleanIndex(last)

	if first >= last {
		return list.Copy()
	}

	var newList AnimalsList
	newList = append(newList, list[0:first]...)
	newList = append(newList, list[last:]...)

	return newList
}

// Part returns a portion of a AnimalsList as a new list.
func (list AnimalsList) Part(first, last int) AnimalsList {
	first = list.cleanIndex(first)
	last = list.cleanIndex(last)

	if first >= last {
		return NewAnimalsList()
	}

	return list[first:last]
}

// Code generated with make-binois-pointer-list.sh DO NOT EDIT.
// http://github.com/pdk/binoislist

package models

import (
	"fmt"
	"sort"
	"strings"
)

// UserList is a slice of *User.
type UserList []*User


// NewUserList creates a new UserList with some initial items.
func NewUserList(items ...*User) UserList {
	return items
}

// Append adds one or more items to the list.
// Don't forget to assign the result to your variable.
//   myList = myList.Append(x, y)
func (list UserList) Append(items ...*User) UserList {
	return append(list, items...)
}

// Cat concatenates multiple lists into a single list.
// Don't forget to assign the result to your variable.
//   myList = myList.Cat(l1, l2)
func (list UserList) Cat(inputLists ...UserList) UserList {
	for _, oneList := range inputLists {
		list = append(list, oneList...)
	}
	return list
}

// Join converts all the items in a UserList to strings and joins them with a
// separator. Default separator is ", ".
func (list UserList) Join(seps ...string) string {
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
func (list UserList) Len() int {
	return len(list)
}

// Get returns the item at position i.
func (list UserList) Get(i int) *User {
	return list[i]
}

// GetOrUnknown returns the element as position i, or unknown if out of range.
func (list UserList) GetOrUnknown(i int) (*User, error) {
	if i < 0 || i >= len(list) {
		return nil, fmt.Errorf("out of range")
	}
	return list[i], nil
}

// UserMapperFunc processes a *User and returns a *User.
type UserMapperFunc func(g *User) *User

func (list UserList) Map(f UserMapperFunc) UserList {
	var newList UserList
	for _, item := range list {
		newList = append(newList, f(item))
	}
	return newList
}

// UserTesterFunc is a function that returns true/false given a *User. Used
// for finding or filtering with a UserList.
type UserTesterFunc func(g *User) bool

// Filter returns a new UserList of items where f() is true.
func (list UserList) Filter(f UserTesterFunc) UserList {
	var newList UserList
	for _, item := range list {
		if f(item) {
			newList = append(newList, item)
		}
	}
	return newList
}

// First returns the first *User where f() is true.
func (list UserList) First(f UserTesterFunc) (item *User, found bool) {
	for _, item := range list {
		if f(item) {
			return item, true
		}
	}

	return nil, false
}

// UserLessThanFunc is a function that compares two items of type *User.
type UserLessThanFunc func(a, b *User) bool

// Sort sorts IN PLACE a UserList with the input comparer function.
func (list UserList) Sort(f UserLessThanFunc) UserList {
	sort.Slice(list, func(i, j int) bool {
		return f(list[i], list[j])
	})

	return list
}

// Copy returns a UserList containing the contents of the list.
func (list UserList) Copy() UserList {
	newList := make(UserList, len(list))
	copy(newList, list)
	return newList
}

// UserEqualizer provides an Equals method for items of type *User.
type UserEqualizer interface {
	Equals(*User) bool
}

// Equalizer handles type conversion. This is a workaround to allow users
// of this package to partially implement requirements.
func (thing *User) Equalizer() UserEqualizer {
	var x interface{} = thing
	return x.(UserEqualizer)
}

// affirmUserEqualsImplemented exists to output a more meaningful error when
// the Equals method has not been implemented. We could have just done type
// assertions w/o checking and had the same net result (panic), but having a
// better explanation is better.
func affirmUserEqualsImplemented(methodName string) {
	var t *User
	var x interface{} = t
	_, ok := x.(UserEqualizer)
	if !ok {
		panic("implement method *User.Equals(*User) in order to use " + methodName)
	}
}

// Index returns the position  of item in the list, and if found.
// Method Equals(*User) must be implemented.
func (list UserList) Index(item *User) (int, bool) {
	affirmUserEqualsImplemented("UserList.Index(item)")

	eqItem := item.Equalizer()

	for pos, item := range list {
		if eqItem.Equals(item) {
			return pos, true
		}
	}

	return -1, false
}

// LastIndex returns the position  of item in the list, and if found.
// Method Equals(*User) must be implemented.
func (list UserList) LastIndex(item *User) (int, bool) {
	affirmUserEqualsImplemented("UserList.LastIndex(item)")

	eqItem := item.Equalizer()

	for pos := len(list) - 1; pos >= 0; pos-- {
		if eqItem.Equals(list[pos]) {
			return pos, true
		}
	}

	return -1, false
}

// Equals checks if two UserList instances are equal by comparing all their
// elements with the *User Equals method.
func (list UserList) Equals(other UserList) bool {
	affirmUserEqualsImplemented("UserList.Equals()")

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
func (list UserList) Delete(index int) UserList {
	if len(list)+index < 0 {
		return list.Copy()
	}
	index = list.cleanIndex(index)
	return list.DeletePart(index, index+1)
}

// cleanIndex is a helper method to ensure offsets are within range.
func (list UserList) cleanIndex(i int) int {
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

// DeletePart returns a new UserList with a portion omitted.
func (list UserList) DeletePart(first, last int) UserList {
	first = list.cleanIndex(first)
	last = list.cleanIndex(last)

	if first >= last {
		return list.Copy()
	}

	var newList UserList
	newList = append(newList, list[0:first]...)
	newList = append(newList, list[last:]...)

	return newList
}

// Part returns a portion of a UserList as a new list.
func (list UserList) Part(first, last int) UserList {
	first = list.cleanIndex(first)
	last = list.cleanIndex(last)

	if first >= last {
		return NewUserList()
	}

	return list[first:last]
}

package models

import (
	"strings"
	"testing"
)

func TestNewUserList(t *testing.T) {
	l := NewUserList()

	u := NewUser("Bob", "bob@nowhere.com")

	l = l.Append(u)

	if l.Len() != 1 {
		t.Errorf("expected l.Len() to return 1, but got %d", l.Len())
	}
}

func TestIndex(t *testing.T) {
	l := NewUserList(
		NewUser("Alice", "alice@wonderland.com"),
		NewUser("Bob", "bob@nowhere.com"),
		NewUser("Charles", "charlie@chocolatefactory.com"),
	)

	i, present := l.Index(NewUser("Bob", "bob@nowhere.com"))

	if !present {
		t.Errorf("expected l.Index() to find something, but failes")
	}

	if i != 1 {
		t.Errorf("expected l.Index() to return 1, but got %d", i)
	}
}

func TestSort(t *testing.T) {
	l := NewUserList(
		NewUser("Charles", "charlie@chocolatefactory.com"),
		NewUser("Bob", "bob@nowhere.com"),
		NewUser("Alice", "alice@wonderland.com"),
	)

	userNameComparator := func(a, b *User) bool {
		return strings.ToLower(a.Name) < strings.ToLower(b.Name)
	}

	l.Sort(userNameComparator)

	if l.Get(0).Name != "Alice" {
		t.Errorf("expected first user after sort to be Alice, but got %s", l.Get(0).Name)
	}
	if l.Get(1).Name != "Bob" {
		t.Errorf("expected second user after sort to be Bob, but got %s", l.Get(1).Name)
	}
	if l.Get(2).Name != "Charles" {
		t.Errorf("expected third user after sort to be Charles, but got %s", l.Get(2).Name)
	}
}

func TestMap(t *testing.T) {
	l := NewUserList(
		NewUser("Alice", "Alice@WonderLand.com"),
		NewUser("Bob", "Bob@NoWhere.com"),
		NewUser("Charles", "Charlie@ChocolateFactory.com"),
	)

	emailToLower := func(user *User) *User {
		return NewUser(user.Name, strings.ToLower(user.Email))
	}

	l = l.Map(emailToLower)

	if l.Get(0).Email != "alice@wonderland.com" {
		t.Errorf("expected email to be alice@wonderland.com, but got %s", l.Get(0).Email)
	}
	if l.Get(1).Email != "bob@nowhere.com" {
		t.Errorf("expected email to be bob@nowhere.com, but got %s", l.Get(1).Email)
	}
	if l.Get(2).Email != "charlie@chocolatefactory.com" {
		t.Errorf("expected email to be charlie@chocolatefactory.com, but got %s", l.Get(2).Email)
	}
}

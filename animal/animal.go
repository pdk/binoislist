package animal

import (
	"fmt"
)

// Animal is a named, (not) furry thing with a color.
type Animal struct {
	Name  string
	Furry bool
	Color string
}

//go:generate $GOPATH/src/github.com/pdk/binoislist/make-binois-list.sh animal_list.go animal Animal AnimalsList Animal{}

func (a Animal) String() string {
	return fmt.Sprintf("Animal{name: %s, furry: %v, color: %s}", a.Name, a.Furry, a.Color)
}

func (a Animal) Equals(b Animal) bool {
	if a.Name == b.Name {
		return true
	}
	return false
}

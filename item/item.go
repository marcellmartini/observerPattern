package item

import (
	"fmt"

	"github.com/marcelma/observerPattern/customer"
	"github.com/marcelma/observerPattern/observer"
)

// Item bla bla bla
type Item struct {
	observer.Observable
	Name string
}

// CreateItem bla bla bla
func CreateItem(name string) *Item {
	return &Item{Name: name}
}

// SetName bla bla bla
func (i *Item) SetName(name string) {
	i.Name = name
	i.NotifyAll(name)
}

func (i *Item) String() string {
	s := "{"

	s += fmt.Sprintf("\"item\": \"%v\", ", i.Name)

	s += fmt.Sprintf("\"observers\": [")
	s += fmt.Sprintf("\"%v\"", i.Observers[0].(*customer.Customer).Name)
	for _, obs := range i.Observers[1:] {
		s += fmt.Sprintf(", \"%v\"", obs.(*customer.Customer).Name)
	}
	s += "]"

	s += "}"

	return s
}

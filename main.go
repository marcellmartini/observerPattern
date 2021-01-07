package main

import (
	"github.com/marcelma/observerPattern/customer"
	"github.com/marcelma/observerPattern/observer"
)

// Item bla bla bla
type Item struct {
	observer.Observable
	Name string
}

func createItem(name string) *Item {
	return &Item{Name: name}
}
func (i *Item) setName(name string) {
	i.Name = name
	i.NotifyAll(name)
}
func main() {
	i := createItem("Nike Shoes")
	c1 := customer.CreateCustomer("Raj")
	c2 := customer.CreateCustomer("Rohit")
	i.AddObserver(c1)
	i.AddObserver(c2)
	i.setName("Adidas")
}

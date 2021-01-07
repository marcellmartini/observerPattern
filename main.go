package main

import (
	"fmt"

	"github.com/marcelma/observerPattern/customer"
	"github.com/marcelma/observerPattern/item"
)

// Item bla bla bla
func main() {
	i := item.CreateItem("Nike Shoes")

	c1 := customer.CreateCustomer("Raj")
	c2 := customer.CreateCustomer("Rohit")
	c3 := customer.CreateCustomer("Marcell")

	i.AddObserver(c1)
	i.AddObserver(c2)
	i.AddObserver(c3)

	i.SetName("Adidas")
	fmt.Println(i)

	i.SetName("Nike")
	fmt.Println(i)
}

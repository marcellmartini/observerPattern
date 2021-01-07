package customer

import "fmt"

// CreateCustomer bla bla bla
func CreateCustomer(name string) *Customer {
	return &Customer{Name: name}
}

// Customer Bla Bla bla
type Customer struct {
	Name string
}

// Update Bla bla bla
func (c *Customer) Update(value string) {
	fmt.Printf("Customer : %s , New value: %s\n", c.Name, value)
}

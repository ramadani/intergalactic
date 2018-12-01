package credit

import (
	"fmt"
)

// Credit has one or more credit types
type Credit struct {
	types []*Type
}

// Type has name and amount attribute
type Type struct {
	name   string
	amount float64
}

// AddType to add a type to credit
func (c *Credit) AddType(name string, amount float64) error {
	tp, err := c.GetType(name)
	if err == nil {
		return fmt.Errorf("%s credit is exists", tp.name)
	}

	c.types = append(c.types, &Type{name, amount})

	return nil
}

// GetType to get a type from credit
func (c *Credit) GetType(name string) (*Type, error) {
	for _, tp := range c.types {
		if name == tp.name {
			return tp, nil
		}
	}

	return nil, fmt.Errorf("%s credit type is not found", name)
}

// NewCredit to make instance of credit
func NewCredit() *Credit {
	return &Credit{}
}

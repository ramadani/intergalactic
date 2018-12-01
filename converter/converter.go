package converter

import (
	"fmt"
	"strings"

	"github.com/ramadani/intergalactic/numeral"
)

// Converter for foreign language to understandable numerals
type Converter struct {
	engine numeral.Engine
	units  []*Unit
}

// Unit for foreign language
type Unit struct {
	alias string
	symb  string
}

// AddUnit for conversion from foreign language
func (c *Converter) AddUnit(alias, symb string) error {
	u, err := c.getUnit(alias)
	if err == nil {
		return fmt.Errorf("%s unit is exists", u.alias)
	}

	c.units = append(c.units, &Unit{alias, symb})

	return nil
}

// GetNum from foreign language based on numeral engine
func (c *Converter) GetNum(alias string) (int, error) {
	aliasArr := strings.Split(alias, " ")
	var symbs []string

	for _, al := range aliasArr {
		u, err := c.getUnit(al)
		if err != nil {
			return 0, err
		}

		symbs = append(symbs, u.symb)
	}

	sSymb := strings.Join(symbs, "")

	return c.engine.ToNumber(sSymb)
}

func (c *Converter) getUnit(alias string) (*Unit, error) {
	for _, u := range c.units {
		if u.alias == alias {
			return u, nil
		}
	}

	return nil, fmt.Errorf("Unit not found")
}

// NewConverter to make instance of converter
func NewConverter(engine numeral.Engine) *Converter {
	return &Converter{engine: engine}
}

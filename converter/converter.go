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
func (c *Converter) AddUnit(alias, symb string) {
	c.units = append(c.units, &Unit{alias, symb})
}

// GetNum from foreign language based on numeral engine
func (c *Converter) GetNum(alias string) (int, error) {
	aliasArr := strings.Split(alias, " ")
	var symbs []string

	for _, al := range aliasArr {
		symb, err := c.getUnitSymbol(al)
		if err != nil {
			return 0, err
		}

		symbs = append(symbs, symb)
	}

	sSymb := strings.Join(symbs, "")

	return c.engine.ToNumber(sSymb)
}

func (c *Converter) getUnitSymbol(alias string) (string, error) {
	for _, u := range c.units {
		if u.alias == alias {
			return u.symb, nil
		}
	}

	return "", fmt.Errorf("Not found")
}

// NewConverter to make instance of converter
func NewConverter(engine numeral.Engine) *Converter {
	return &Converter{engine: engine}
}

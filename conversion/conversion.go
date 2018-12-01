package conversion

import (
	"fmt"
	"strings"

	"github.com/ramadani/intergalactic/converter"
	"github.com/ramadani/intergalactic/credit"
	"github.com/ramadani/intergalactic/querier"
)

// Conversion a query of foreign language to understandable language
type Conversion struct {
	querier   *querier.Querier
	converter *converter.Converter
	credit    *credit.Credit
}

// Query is method for conversion a query of foreign language to understandable language
func (c *Conversion) Query(stmt string) (string, error) {
	stmtArr := strings.Split(stmt, " ")

	if c.querier.IsType1(stmtArr) {
		alias, symbol := c.querier.GetType1Values(stmtArr)
		err := c.converter.AddUnit(alias, symbol)
		if err != nil {
			return "", err
		}

		return "", nil
	} else if c.querier.IsType2(stmtArr) {
		alias, creditTypeName, total := c.querier.GetType2Values(stmtArr)
		num, err := c.converter.GetNum(alias)
		if err != nil {
			return "", err
		}

		amount := float64(total) / float64(num)
		err = c.credit.AddType(creditTypeName, amount)
		if err != nil {
			return "", err
		}

		return "", nil
	} else if c.querier.IsType3(stmtArr) {
		alias := c.querier.GetType3Values(stmtArr)
		num, err := c.converter.GetNum(alias)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s is %d", alias, num), nil
	} else if c.querier.IsType4(stmtArr) {
		alias, creditTypeName := c.querier.GetType4Values(stmtArr)

		num, err := c.converter.GetNum(alias)
		if err != nil {
			return "", err
		}

		creditType, err := c.credit.GetType(creditTypeName)
		if err != nil {
			return "", err
		}

		total := float64(num) * creditType.Amount

		return fmt.Sprintf("%s %s is %.0f Credits", alias, creditTypeName, total), nil
	}

	return "I have no idea what you are talking about", nil
}

// NewConversion to make instance of Conversion
func NewConversion(querier *querier.Querier, converter *converter.Converter,
	credit *credit.Credit) *Conversion {
	return &Conversion{querier, converter, credit}
}

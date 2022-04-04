package domain

import (
	"fmt"
)

func TestingErrorNotMatched(number uint8, expected interface{}, actual interface{}) string {
	return fmt.Sprintf("\ntest#%d\nexpect = %#v \n----\nactual = %#v", number, expected, actual)
}

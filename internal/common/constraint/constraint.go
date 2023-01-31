package constraint

import (
	"fmt"
	"reflect"
)

func NotNil(paramName string, value any) {
	if reflect.ValueOf(value).IsNil() { // == no workie
		panic("Function parameter '" + paramName + "' should not be nil")
	}
}

func NotEq[T comparable](paramName string, value T, shouldNotBe T) {
	if value == shouldNotBe {
		panic(fmt.Sprintf("Function parameter '"+paramName+"' should not be %s", shouldNotBe))
	}
}

package arrays_test

import (
	"go-service-echo/util/arrays"
	"testing"
)

func TestInArray(t *testing.T) {
	list := []string{"1", "2", "3", "4", "5"}
	act := arrays.InArray(list, "2")
	act2 := arrays.InArray(list, "0")

	if true != act {
		t.Error("must same")
	}

	if false != act2 {
		t.Error("must same")
	}
}

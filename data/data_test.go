package data_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/vickz86/vegePrice/data"
)

var testVege = data.Vege{
	Name:     "spinash",
	PrizeVar: []float64{1.0, 2.0, 3.0},
}

func TestAddNewPrice(t *testing.T) {
	want := []float64{1.0, 2.0, 3.0, 4.0}

	//add to vege prize
	testVege.AddNewPrice(4.0)

	//compare slices
	compVal := slices.Compare(want, testVege.PrizeVar)

	if compVal != 0 {
		fmt.Printf("want %v and got %v", want, testVege.PrizeVar)

	}

}

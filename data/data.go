package data

import (
	"fmt"
	"log"

	mathuti "github.com/vickz86/GoFunctions/mathUti"
)

// declare the type
type Vege struct {
	Name     string
	PrizeVar []float64
	Amount   int
	TotalVal float64
}

// create a slice of type Vege
var VegeItem = []Vege{
	{
		Name:     "carrot",
		PrizeVar: []float64{1.5, 1.8, 1.4},
		Amount:   10,
	},
	{
		Name:     "potato",
		PrizeVar: []float64{0.9, 1.1, 1.25},
		Amount:   30,
	},
	{
		Name:     "zuchini",
		PrizeVar: []float64{2.2, 2.8, 2.7, 3.1, 3.5},
		Amount:   30,
	},
}

// PrintCurPrize : print current prize of vegetable (last value in slice)
func (v *Vege) PrintCurPrize() {
	//currentPrice
	current := v.PrizeVar[len(v.PrizeVar)-1]

	fmt.Printf("the price of %s is %.1f\n", v.Name, current)

}

// print and return percent variation from the original price
func (v *Vege) PercVarBeginning() float64 {
	varia, err := mathuti.PercVarWithLastAndOther(v.PrizeVar, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(varia)

	return varia

}

func (v *Vege) AddNewPrice(toAdd float64) {
	v.PrizeVar = append(v.PrizeVar, toAdd)
	// TODO , write test
}

package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/vickz86/vegePrice/data"
)

var StringToDo = "type to do\n0 exit\n1 print current prize of all vegetable\n" +
	"2 print current price of specific vegetable\n3 add new prize to vegetable"

func main() {
	var toDo int
	// loop action
	for {
		fmt.Println(StringToDo)
		fmt.Scan(&toDo)
		//switch
		switch toDo {
		case 0:
			os.Exit(1)
		case 1:
			for _, vege := range data.VegeItem {
				vege.PrintCurPrize()

			}
		case 2:
			theVeg := VegeFromList(data.VegeItem)
			theVeg.PrintCurPrize()
		case 3:
			vegeAddPriz := VegeFromList(data.VegeItem)
			newPrize := NewPrizeForVege(vegeAddPriz)
			//create an instance with new prize
			vegeAddPriz.AddNewPrice(newPrize)
			data.VegeItem = AddToVegePrize(data.VegeItem, vegeAddPriz)
		}

	}

}

// ----function----

// VegeFromList : return a type of Vege from slice of vege based on name
func VegeFromList(sliceVege []data.Vege) data.Vege {
	var theName string

	//loop find vege
	for {
		fmt.Println("type vege:")
		fmt.Scan(&theName)

		for _, veg := range sliceVege {
			if veg.Name == theName {
				return veg

			}

		}

		fmt.Println("couldnt find your vege , try again , here are list of vege")

		for _, veg := range sliceVege {
			fmt.Println(veg.Name)

		}

	}

}

// utility function
// TODO Add new price to a vege
func NewPrizeForVege(theVege data.Vege) float64 {
	//variable
	var newPrize float64

	//ask new prize
	fmt.Printf("type new prize of  %s , last prize was %.1f\n", theVege.Name, theVege.PrizeVar[len(theVege.PrizeVar)-1])
	fmt.Scan(&newPrize)

	return newPrize

}

// add instance with newPrize to slice of vege , return a slice of Vege with new Prize
func AddToVegePrize(sliceVege []data.Vege, vegeToAdd data.Vege) []data.Vege {
	//get the index of vege you want to change
	var indexVege int

	for index, vege := range sliceVege {
		if vege.Name == vegeToAdd.Name {
			indexVege = index
		}

	}

	//replace in slice vege
	slices.Replace(sliceVege, indexVege, indexVege+1, vegeToAdd)

	//return
	return sliceVege

}

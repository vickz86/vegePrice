package main

import (
	"fmt"
	"os"

	"github.com/vickz86/vegePrice/data"
)

var StringToDo = "type to do\n0 exit\n1 print current prize of all vegetable\n" +
	"2 print current price of specific vegetable"

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

// TODO Add new price to a vege

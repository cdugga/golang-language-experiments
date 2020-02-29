package main

import "fmt"

func testSwitch(cryptoCoin string) {
	switch coin := cryptoCoin; coin{
	case "NEO":
		fmt.Println("Neo currency")
	case "ZCASH":
		fmt.Println("Zcash ZPK currency")
	case "BTC":
		fmt.Println("Bitcoin currency")
	default:
		fmt.Println("%s.\n", coin)
	}
}

func testTypeSwitch(){

	var a FullStack
	// a type implements an interface implicitly by implementing its methods
	a = dev{
		name: "colin",
		email: contactAddr{
			email:"cd@gmail.com",
		},
		proficientIn: languages{
			[]string{"Java", "Python", "Kotlin", "GO"},
		},
	}


	switch v := a.(type) {
	case dev:
		fmt.Println("Type Dev", v)
	default:
		fmt.Println("%s.\n", "No type match found", v)
	}
}

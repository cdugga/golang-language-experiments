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

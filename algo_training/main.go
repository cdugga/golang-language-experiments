package main

import (
	"fmt"
	"strings"
)




func main(){
	var results = make([]string,4)
	results[0] = ToCamelCase("The_Stealth_Warrior")
	results[1] = ToCamelCase("The-Stealthy-Warrior")
	results[2] = ToCamelCase("the_Stealth_Warrior")
	results[3] = ToCamelCase("the-Stealthy-Warrior")
	results[3] = ToCamelCase("up_bridge_side_south_Samurai_bridge")

	for _, result := range results{
		fmt.Println(result)
	}
}

func ToCamelCase(s string) string {
	if s == "" {return s}
	t := strings.Split(strings.Replace(strings.Replace(s, "_", " ", -1), "-"," ", -1), " ")

	if t[0][:1] != strings.ToUpper(s[:1]) {
		return strings.Replace(t[0] + strings.Title(strings.Join(t[1:], " ")), " ", "",-1)
	} else {
		return strings.Replace(strings.Title(strings.Join(t[0:]," ")), " ", "",-1)
	}
}













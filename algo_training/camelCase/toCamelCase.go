package camelCase

import (
	"fmt"
	"strings"
)

func TestCamelCase(){
	var results = make([]string,5)
	results[0] = toCamelCase("The_Stealth_Warrior")
	results[1] = toCamelCase("The-Stealthy-Warrior")
	results[2] = toCamelCase("the_Stealth_Warrior")
	results[3] = toCamelCase("the-Stealthy-Warrior")
	results[4] = toCamelCase("up_bridge_side_south_Samurai_bridge")

	for _, result := range results{
		fmt.Println(result)
	}
}


func toCamelCase(s string) string {
	if s == "" {return s}
	t := strings.Split(strings.Replace(strings.Replace(s, "_", " ", -1), "-"," ", -1), " ")

	if t[0][:1] != strings.ToUpper(s[:1]) {
		return strings.Replace(t[0] + strings.Title(strings.Join(t[1:], " ")), " ", "",-1)
	} else {
		return strings.Replace(strings.Title(strings.Join(t[0:]," ")), " ", "",-1)
	}
}
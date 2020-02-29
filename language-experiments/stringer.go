package main

import "fmt"

type developer struct {
	name string
	company string
	title string
	primaryLangauage string
	frameworks []string
}

func (d developer) String() string {
	return fmt.Sprintf("%v (%v LTD)", d.name, d.company, d.title, d.primaryLangauage)
}

func testStringer(){
	fmt.Println(developer{
		"colin",
		"bridgemate",
		"tech lead",
		"java",
		[]string{
			"Kafka",
			"Kubernetes"}})
}

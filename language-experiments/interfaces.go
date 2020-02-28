package main

import "fmt"

type FullStack interface{
	languages() []string
}

func (d dev) languages() []string{
	return d.proficientIn.languages
}

func testInterface(){

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
	fmt.Println(a.languages())

}
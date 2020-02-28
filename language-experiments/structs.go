package main

import "strings"

type contactAddr struct {
	email string
}

type languages struct {
	languages []string
}

type dev struct {
	name string
	email contactAddr
	proficientIn languages
}

// will update reference to pointer
func (d dev) toUpper() (string, string) {
	return strings.ToUpper(d.name), strings.ToUpper(d.email.email)
}



func createDev() dev {
	d := dev{
		name: "Colin",
		email: contactAddr{
			email: "duggan.colin@gmail.com",
		},
	}
	return d
}
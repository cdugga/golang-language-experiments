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

// will update the point in memory
func (d *dev) toUpperWithPointer() {
	(*d).email.email = strings.ToUpper(d.email.email)
	(*d).name = strings.ToUpper(d.name)
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
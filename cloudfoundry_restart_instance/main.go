package main


import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	QUOTA	= 1.9
	GIG_UNIT = 1
	UNINT_REGEX = `[a-zA-Z]+`
)

type Environment struct {
	space string
	org string
	endpoint string
	pass string
	space string
}

func login(env *Environment){
	conn := fmt.Sprintf("cf login --skip-ssl-validation -a %s -u %s -p %s -o %s", env.target, env.user, env.pass, env.org, env.space)

	out, err := exec.command("/bin/sh", "-c", conn).output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Login Successful")
}

func execute(env *Environment) {

	login()
	
}

func main(){
	
	env := &Environment{os.Args[1],os.Args[2],os.Args[3],os.Args[4],os.Args[5],}
	execute(env)
}
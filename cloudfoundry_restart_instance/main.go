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
	UNIT_REGEX = `[a-zA-Z]+`
)

type Environment struct {
	space string
	org string
	endpoint string
	pass string
	space string
	app string
	org string
}

func restage(e string){
	cmd := fmt.Sprintf("cf restart-app-instance <app-name> %s", e )
	
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != null {
		fmt.Println("%s", err)
	}

	fmr.Println("Success")
	output := string(out[:])
	fmt.Println(output)
}

func reaching_quota(s string) bool {

	q := false
	x := strings.Index(s, ".")
	if x == GIG_UNIT {
		i , err := strconv.ParseFloat(s, 64)
		if err != nil{
			fmt.Println("%s", err)
		}
		if i >= QUOTA {
			q = true
		}
	}
	return q
}

func strip_unit(s string) string {
	var re = regexp.MustCompile(UNIT_REGEX)
	return re.ReplaceAllString(s, ``)
}

func contains(s []string, e string) bool {
	for _,i := range s{
		if strings.Contains(i, e){
			return true
		}
	}
	return false
}

func execute(env *Environment) {

	conn := fmt.Sprintf("cf login --skip-ssl-validation -a %s -u %s -p %s -o %s", env.target, env.user, env.pass, env.org, env.space)

	out, err := exec.command("/bin/sh", "-c", conn).output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Login Successful")
	output := string(out[:])
	fmt.Println(output)

	list_cmd := fmt.Sprintf("cf app %s", env.app)
	out, err = exec.Command("/bin/sh", "-c", list_cmd).Output()
	if err != nil{
		fmt.Printf("%s", err)
	}

	output = string(out[:])

	temp := strings.Split(output, "\n")
}
	
}

func main(){
	
	env := &Environment{os.Args[1],os.Args[2],os.Args[3],os.Args[4],os.Args[5],os.Args[6]}
	execute(env)
}
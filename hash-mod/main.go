package main

import (

	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

func fetch(){

	endpoint := "https://bing.com/covid/data"
	resp, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Printf(responseString)

}


func main() {

	var name,language string
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "lang, l",
						Value: "english",
						Usage: "Language for the greeting",
						Destination: &language,
					},
					&cli.StringFlag{
							Name:  "name, n",
						Usage: "Enter name",
						Destination: &name,
					},
				},
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "print statistics",

				Action: func(c *cli.Context) error {
					//name := "Nefertiti"
					fetch()
					if c.NArg() > 0 {
						name = c.Args().Get(0)
					}
					if language == "spanish" {
						fmt.Println("Hola", name)
					} else {
						fmt.Println("Hello", name)
					}
					return nil
				},
			},

		},

	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
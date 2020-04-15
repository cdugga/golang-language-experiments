package main

import (
	"encoding/json"
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
		log.Fatal(err)//
	}

	var data Covid19

	json.Unmarshal(responseData, &data)

	fmt.Printf("Results: %v\n", data)
	//responseString := string(responseData)
	//fmt.Printf(responseString)

}

func main() {

	var name,language string
	app := &cli.App{
		Name: "Apigee Publish API CLI Utility",
		//Usage: "Create API Proxy Endpoints",
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "lang, l",
						Aliases: []string{"l"},
						Value: "english",
						Usage: "Language for the greeting",
						//Destination: &language,
					},
					&cli.StringFlag{
						Name:  "name, n",
						Aliases: []string{"n"},
						Usage: "Enter name",
						//Destination: &name,
					},
				},
				Name:    "proxy",
				Aliases: []string{"p"},
				Usage:   "Create Apigee API Proxy",

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
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "lang, l",
						Aliases: []string{"l"},
						Value: "english",
						Usage: "Language for the greeting",
						Destination: &language,
					},
					&cli.StringFlag{
						Name:  "name, n",
						Aliases: []string{"n"},
						Usage: "Enter name",
						Destination: &name,
					},
				},
				Name:    "print",
				Aliases: []string{"pr"},
				Usage:   "Create Apigee Product",

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
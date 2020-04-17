package main

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

const (
	covid19Endpoint = "https://bing.com/covid/data"
)

func find (a [] Areas, x string) int {
	for i, n := range a {
		if x == n.ID {
			return i;
		}
	}
	return len(a)
}

func fetch(z string, d bool) {
	endpoint := covid19Endpoint
	resp, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Covid19
	json.Unmarshal(responseData, &data)
	a := data.getCountry(z)

	if d {
		a.printNested()
		return
	}
	a.print()
}

func main() {

	var name,country string
	app := &cli.App{
		Name: "Latest Covid19 Statistics from around the world.",
		Usage: "Fetch latest published statistics per country",
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "country, c",
						Aliases: []string{"c"},
						Value: "ireland",
						Usage: "Country (without spaces E.g. unitedstates, unitedkingdom, ireland)",
						Destination: &country,
					},
				},
				Name:    "stats",
				Aliases: []string{"s"},
				Usage:   "Display latest Covid19 statistics from around the world.",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						name = c.Args().Get(0)
					}
					fetch(country, false)
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "detailed",
						Usage: "Display detailed regional information (if published).",
						Action: func(c *cli.Context) error {
							//fmt.Println("something new: ", c.Args().First())
							fetch(country, true)
							return nil
						},
					},
				},
			},
		},
	}

	//sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
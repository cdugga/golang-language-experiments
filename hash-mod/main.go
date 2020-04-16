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


func (a *Areas) print() {
	fmt.Printf("Country: %v\n", a.DisplayName);
	fmt.Printf("Last Updated: %v\n", a.LastUpdated);
	fmt.Printf("Total Confirmed Cases: %v\n", a.TotalConfirmed);
	fmt.Printf("Total Deaths: %v\n", a.TotalDeaths);
	fmt.Printf("Total Recovered: %v\n", a.TotalRecovered);
	fmt.Printf("Total Confirmed Delta: %v\n", a.TotalConfirmedDelta);
	fmt.Printf("Total Recovered Delta: %v\n", a.TotalRecoveredDelta);
}

func find (a [] Areas, x string) int {
	for i, n := range a {
		if x == n.ID {
			return i;
		}
	}
	return len(a)
}

func (c *Covid19)  areaDetails (p int) Areas {
		return c.Areas[p]
}

func (c *Covid19) getCountry(x string) Areas {
	i := find(c.Areas, x)
	return c.areaDetails(i)
}

func fetch(z string) {
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
	a := data.getCountry(z)
	a.print()
}

func main() {

	var name,country string
	app := &cli.App{
		Name: "Covid19 Latest Statistics",
		Usage: "Search for latest statistics based on country name (no spaces)",
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "country, c",
						Aliases: []string{"c"},
						Value: "ireland",
						Usage: "Country (without spaces E.g. unitedstates)",
						Destination: &country,
					},
				},
				Name:    "stats",
				Aliases: []string{"s"},
				Usage:   "Create Apigee API Proxy",

				Action: func(c *cli.Context) error {

					if c.NArg() > 0 {
						name = c.Args().Get(0)
					}
					fetch(country)

					return nil
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
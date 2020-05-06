package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"sort"
)


type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type CloudDataType int

const (
	Resource = iota
	Pricing
)


func serveDynamicResults(c CloudType) {

	http.HandleFunc("/inventory", c.ShowHandler)
	http.ListenAndServe(":3000", nil)
}


func fetchData (c CloudType, r string, dataType CloudDataType){

	switch dataType {
	case Pricing: {
		c.GetCostAndUsage()
	}
	case Resource: {
		c.GetResources()

		serveDynamicResults(c)
	}
	default: {
		fmt.Println("Missing data type ")
	}
	}
}

func fetchResources(p string, r string, dt CloudDataType) {
	switch p {
	case "aws":	{
			aws := &AWSCloud{
				region: "us-east-1",
				limit:  "5",
			}
			fetchData(aws, r, dt)
		}
	default:
		{
			fmt.Printf("Please choose from on of the supported cloud platforms [AWS|GCP|AZURE]")
		}
	}
}

func main() {
	var provider, resource string = "aws", "all"
	var dataType CloudDataType = Resource

	myFigure := figure.NewFigure("CloudList", "speed", true)
	myFigure.Print()

	fmt.Println("....")
	fmt.Println("....")
	fmt.Println("Output available on http://localhost:3000/inventory ")

	app := &cli.App{
		Name: "Cloud Automation Tool.",
		Usage: "Fetch Resource usage and pricing across AWS, GCP and Azure cloud providers",
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "provider, p",
						Aliases: []string{"p"},
						Value: "aws",
						Usage: "Supply cloud provider [AWS|GCP|Azure]",
						Destination: &provider,
					},
				},
				Name:    "resources",
				Aliases: []string{"r"},
				Usage:   "Lists your cloud resources.",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						provider = c.Args().Get(0)
					}
					fetchResources(provider, resource, dataType)
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "service-code",
						Usage: "Lists specific resources based on supplied value. [AmazonS3 etc.]",
						Action: func(c *cli.Context) error {
							//fmt.Println("something new: ", c.Args().First())
							if c.NArg() > 0 {
								resource = c.Args().Get(0)
							}
							fetchResources(provider, resource, dataType)
							return nil
						},
					},
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "provider, p",
						Aliases: []string{"p"},
						Value: "aws",
						Usage: "Supply cloud provider [AWS|GCP|Azure]",
						Destination: &provider,
					},
				},
				Name:    "pricing",
				Aliases: []string{"p"},
				Usage:   "Fetch Usage and Pricing information.",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						provider = c.Args().Get(0)
					}
					dataType = Pricing
					fetchResources(provider, resource, dataType)
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


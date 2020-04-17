package main

import (
	"fmt"
	"time"
)

type area interface {
	print()
	printNested()
}

type covid interface {
	areaDetails() Areas
	getCountry() Areas
}

type Covid19 struct {
	ID                  string    `json:"id"`
	DisplayName         string    `json:"displayName"`
	Areas               []Areas   `json:"areas"`
	TotalConfirmed      int       `json:"totalConfirmed"`
	TotalDeaths         int       `json:"totalDeaths"`
	TotalRecovered      int       `json:"totalRecovered"`
	TotalRecoveredDelta int       `json:"totalRecoveredDelta"`
	TotalDeathsDelta    int       `json:"totalDeathsDelta"`
	TotalConfirmedDelta int       `json:"totalConfirmedDelta"`
	LastUpdated         time.Time `json:"lastUpdated"`
}
type Areas struct {
	ID                  string    `json:"id"`
	DisplayName         string    `json:"displayName"`
	Areas               []Areas   `json:"areas"`
	TotalConfirmed      int       `json:"totalConfirmed"`
	TotalDeaths         int       `json:"totalDeaths"`
	TotalRecovered      int       `json:"totalRecovered"`
	TotalRecoveredDelta int       `json:"totalRecoveredDelta"`
	TotalDeathsDelta    int       `json:"totalDeathsDelta"`
	TotalConfirmedDelta int       `json:"totalConfirmedDelta"`
	LastUpdated         time.Time `json:"lastUpdated"`
	Lat                 float64   `json:"lat"`
	Long                float64   `json:"long"`
	ParentID            string    `json:"parentId"`
}


func (c *Covid19) areaDetails (p int) Areas {
	return c.Areas[p]
}

func (c *Covid19) getCountry(x string) Areas {
	i := find(c.Areas, x)
	return c.areaDetails(i)
}

func (a *Areas) print() {
	fmt.Printf("Country/Region: %v\n", a.DisplayName);
	fmt.Printf("Last Updated: %v\n", a.LastUpdated);
	fmt.Printf("Total Confirmed Cases: %v\n", a.TotalConfirmed);
	fmt.Printf("Total Deaths: %v\n", a.TotalDeaths);
	fmt.Printf("Total Recovered: %v\n", a.TotalRecovered);
	fmt.Printf("Total Confirmed Delta: %v\n", a.TotalConfirmedDelta);
	fmt.Printf("Total Recovered Delta: %v\n", a.TotalRecoveredDelta);
}

func (a *Areas) printNested() {
	fmt.Printf("Country/Region: %v\n", a.DisplayName);
	fmt.Printf("Last Updated: %v\n", a.LastUpdated);
	fmt.Printf("Total Confirmed Cases: %v\n", a.TotalConfirmed);
	fmt.Printf("Total Deaths: %v\n", a.TotalDeaths);
	fmt.Printf("Total Recovered: %v\n", a.TotalRecovered);
	fmt.Printf("Total Confirmed Delta: %v\n", a.TotalConfirmedDelta);
	fmt.Printf("Total Recovered Delta: %v\n", a.TotalRecoveredDelta);

	if len(a.Areas) > 0 {
		l := a.Areas
		for _, r := range l {
			r.print()
		}
	}
}
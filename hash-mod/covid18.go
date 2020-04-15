package main

import "time"

type Covid19 struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Areas       []struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Areas       []struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
			Areas       []struct {
				ID                  string        `json:"id"`
				DisplayName         string        `json:"displayName"`
				Areas               []interface{} `json:"areas"`
				TotalConfirmed      int           `json:"totalConfirmed"`
				TotalDeaths         int           `json:"totalDeaths"`
				TotalRecovered      int           `json:"totalRecovered"`
				TotalRecoveredDelta interface{}   `json:"totalRecoveredDelta"`
				TotalDeathsDelta    int           `json:"totalDeathsDelta"`
				TotalConfirmedDelta int           `json:"totalConfirmedDelta"`
				LastUpdated         time.Time     `json:"lastUpdated"`
				Lat                 float64       `json:"lat"`
				Long                float64       `json:"long"`
				ParentID            string        `json:"parentId"`
			} `json:"areas"`
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
		} `json:"areas"`
		TotalConfirmed      int       `json:"totalConfirmed"`
		TotalDeaths         int       `json:"totalDeaths"`
		TotalRecovered      int       `json:"totalRecovered"`
		TotalRecoveredDelta int       `json:"totalRecoveredDelta"`
		TotalDeathsDelta    int       `json:"totalDeathsDelta"`
		TotalConfirmedDelta int       `json:"totalConfirmedDelta"`
		LastUpdated         time.Time `json:"lastUpdated,omitempty"`
		Lat                 float64   `json:"lat"`
		Long                float64   `json:"long"`
		ParentID            string    `json:"parentId"`
	} `json:"areas"`
	TotalConfirmed      int       `json:"totalConfirmed"`
	TotalDeaths         int       `json:"totalDeaths"`
	TotalRecovered      int       `json:"totalRecovered"`
	TotalRecoveredDelta int       `json:"totalRecoveredDelta"`
	TotalDeathsDelta    int       `json:"totalDeathsDelta"`
	TotalConfirmedDelta int       `json:"totalConfirmedDelta"`
	LastUpdated         time.Time `json:"lastUpdated"`
}
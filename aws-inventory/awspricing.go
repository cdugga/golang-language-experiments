package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func (a *AWSCloud) GetCostAndUsage(){

	s := a.CreateSession()
	sess := s.(*session.Session)
	explorer := costexplorer.New(sess)

	start := "2020-03-01"
	end := "2020-04-01"
	granularity := "MONTHLY"
	var metrics []*string
	costType := "NetUnblendedCost"
	metrics = append(metrics, &costType)

	input := costexplorer.GetCostAndUsageInput{
		Filter:        nil,
		Granularity:   &granularity,
		GroupBy:       nil,
		Metrics:       metrics,
		NextPageToken: nil,
		TimePeriod:    &costexplorer.DateInterval{
			End:   &end,
			Start: &start,
		},
	}

	result, err := explorer.GetCostAndUsage(&input)
	if err == nil {
		fmt.Println(result)
	}
	//a.handleError("Unable to list Cost & Usage Metrics", err)
	//a.printHeader(printSubHeader, "Costs & Usage:")
	for _, b := range result.GroupDefinitions {
		fmt.Printf("* %s %s\n",
			aws.StringValue(b.Key), aws.StringValue(b.Type))
	}
}
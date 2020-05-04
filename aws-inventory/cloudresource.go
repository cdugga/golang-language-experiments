package main

import "net/http"

type CloudType interface {
	CreateSession()interface{}
	GetResources()
	GetCostAndUsage()
	ShowHandler(w http.ResponseWriter, r *http.Request)
}

type CloudResources interface {
	createSession()
	ListStorageBuckets(interface{}) []Bucket
	ListLambda(interface{})
	ListDBTables(interface{})
	ListApiGatewayEndpoints(interface{})
	ListCloudFormationStackSets(interface{})
	ListCloudFormationStack(interface{})
	ListCloudFront(interface{})
}

type CloudPricing interface {
	ListMetrics(interface{})
}


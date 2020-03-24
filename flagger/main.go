package main

import (
    "flag"
    "fmt"
)

func main() {

	swaggerPtr := flag.String("swagger_uri", "", "Swagger.json endpoint")
	apiNamePtr := flag.String("api_name", "", "API Proxy name")
	basePathPtr := flag.String("base_path", "", "Base Path")
	displayPtr := flag.String("api_display_name", "", "API Display Name")
	envPtr := flag.String("env", "", "DEV|QA|PROD")
	descPtr := flag.String("api_desc", "", "API Description")
	intExtPtr := flag.String("internal_external", "", "Internal|External")
	pprfTargetPtr := flag.String("pperf_target", "", "PPERF Target API")
	devTargetPtr := flag.String("dev_target", "", "DEV Target API")
	userPtr := flag.String("apigee_user", "", "Username")
	passPtr := flag.String("apigee_pass", "", "Password")

   
    flag.Parse()

    fmt.Println("swaggerPtr:", *swaggerPtr)
    fmt.Println("apiNamePtr:", *apiNamePtr)
    fmt.Println("basePathPtr:", *basePathPtr)
	fmt.Println("displayPtr:", *displayPtr)
	fmt.Println("envPtr:", *envPtr)
	fmt.Println("descPtr:", *descPtr)
	fmt.Println("intExtPtr:", *intExtPtr)
	fmt.Println("pprfTargetPtr:", *pprfTargetPtr)
	fmt.Println("devTargetPtr:", *devTargetPtr)
	fmt.Println("userPtr:", *userPtr)
	fmt.Println("passPtr:", *passPtr)

}
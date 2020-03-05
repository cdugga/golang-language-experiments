package main

import "language-experiments/request_with_headers"

func main() {

	//httpRequest("http://www.google.com")
	//
	//testPointerUpdate()
	//
	//testForLoop()
	//
	//ifExpr()
	//// implicitly builds array then slices it
	////////////////////////////////////////////
	//coins := []string{"NEO", "ZCASH", "BTC"}
	//testSwitch(coins[0])
	////////////////////////////////////////////
	//testDefer()
	//arrayAndSliceOps()
	//testInterface()

	//testTypeAssertions()
	//testTypeSwitch()
	//stringer.TestStringer()
	//testError()
	//goroutines.TestGoRoutines()

	//http_server.RunServer()
	//api_key_request.FetchEvents()
	//jsonDecoding.TestJson()
	request_with_headers.FetchEventsWithHeaders()
}

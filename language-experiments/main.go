package main

import http_server "language-experiments/http-server"

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

	http_server.RunServer()
}

package main

import (
	rest "xqledger/apirouter/rest"
)

const componentMsg = "API Router"



func main() {
	rest.StartHTTPServer()
}

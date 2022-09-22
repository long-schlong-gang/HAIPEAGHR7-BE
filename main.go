/*
 *
 *	HAIPEAGHR7 Server
 *
 */

package main

import (
	"HAIPEAGHR7/api"
	"HAIPEAGHR7/util"
	"github.com/labstack/echo/v4"
)

func main() {

	// Initialization
	util.LoadConfig()
	util.InitLogger()
	e := echo.New()
	api.InitRoutes(e)

	// Start server
	api.StartServer(e)

	// Termination
	util.TerminateLogger()
}

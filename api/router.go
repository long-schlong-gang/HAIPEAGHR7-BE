/*
 *	Assigns all the handlers to their endpoints
 */

package api

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {

	e.GET("/", helloHandler)

}

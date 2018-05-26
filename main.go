// Develom Website Template
// This go application creates a simple but extremely fast website template using the gin-gonic
// and http packages.
// https://github.com/demo-apps/go-gin-app
// If not specified, the default port is 8080. See router.Run() below.
// !!! Before start using this app, make sure you get the gingonic package
// go get github.com/gin-gonic/gin
// Make sure the Run configuration setup defines a Directory and not a File.
// If it is not set correctly, I will get undefined functions.
// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
// Hector: Since we don't know when we could use the json or xml format and it really doesn't
// hurt to have the render function otherwise, we'll keep the following case statement for now.
// Note: The render function is called by handlers.article.go
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

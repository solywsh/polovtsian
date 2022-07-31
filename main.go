// this file is used to test modules
package main

import (
	"github.com/solywsh/polovtsian/pin"
	"net/http"
)

func main() {
	r := pin.New()

	r.GET("/", func(c *pin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *pin.Context) {
		// expect /hello?name=solywsh
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *pin.Context) {
		// expect /hello/solywsh
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *pin.Context) {
		c.JSON(http.StatusOK, pin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")

}

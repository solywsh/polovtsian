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
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *pin.Context) {
		c.JSON(http.StatusOK, pin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}

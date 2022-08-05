package main

import (
	"github.com/solywsh/polovtsian/pin"
	"net/http"
)

func main() {
	r := pin.Default()
	r.GET("/", func(c *pin.Context) {
		c.String(http.StatusOK, "Hello solywsh\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *pin.Context) {
		names := []string{"solywsh"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}

package main

import (
	"net/http"
	"sloe"
)

func main() {
	r := sloe.Default()
	r.GET("/hello", func(c *sloe.Context) {
		c.String(http.StatusOK, "Hello Web Sole\n")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/get-User-info", func(c *sloe.Context) {
			c.JSON(http.StatusOK, sloe.H{
				"name": "test",
			})
		})
	}

	r.Run(":9999")
}

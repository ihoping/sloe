package main

import (
	"net/http"
	"webSloe"
)

func main() {
	r := webSole.Default()
	r.GET("/hello", func(c *webSole.Context) {
		c.String(http.StatusOK, "Hello Web Sole\n")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/get-User-info", func(c *webSole.Context) {
			c.JSON(http.StatusOK, webSole.H{
				"name": "test",
			})
		})
	}

	r.Run(":9999")
}

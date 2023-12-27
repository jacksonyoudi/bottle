package main

import (
	"bottle/bolt"
	"net/http"
)

func main() {
	r := bolt.New()
	r.GET("/", func(c *bolt.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *bolt.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *bolt.Context) {
		c.JSON(http.StatusOK, bolt.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")

}

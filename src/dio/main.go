package main

import "vendor/github.com/sheng/air"

func main() {
	a := air.New()
	a.GET("/", homeHandler)
	a.Serve()
}

func homeHandler(c *air.Context) error {
	return c.String("Hello, 世界")
}

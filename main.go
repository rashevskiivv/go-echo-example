package main

import (
	"github.com/rashevsiivv/echo-example/router"
)

func main() {
	// create a new echo instance
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}

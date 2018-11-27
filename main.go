package main

import (
	"fmt"

	"github.com/akamiko/ddd-sample/route"
)

func main() {
	fmt.Println("aaa")
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}

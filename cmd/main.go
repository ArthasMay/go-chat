package main

import (
	"fmt"
	"gochat/connect"
)

func main()  {
	connect.New().Run()
}

func test() {
	defer func() {
		fmt.Println("func end")
	}()

	i := 0
	for {
		i ++
		fmt.Println(i)
		if i == 10 {
			return
		}
	}
}
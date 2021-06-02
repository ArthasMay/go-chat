package main

import (
	"fmt"
)

func main()  {
	//var b int64 = 32
	//num := atomic.AddInt64(&b, 1)
	//fmt.Println(num, b)
	test()
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
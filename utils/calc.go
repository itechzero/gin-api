package utils

import "fmt"

type calc func(int, int) int

func doCalc(x, y int, callback calc) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("exception")
	return callback(x, y) //函数作为返回值
}

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

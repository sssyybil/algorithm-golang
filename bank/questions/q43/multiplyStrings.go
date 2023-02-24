package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	n1, err := strconv.ParseInt(num1, 10, 64)
	if err != nil {
		panic(err)
	}
	var ans int64
	ans, zero := 0, 0
	for i := len(num2) - 1; i >= 0; i-- {
		i1, err := strconv.ParseInt(string(num2[i]), 10, 64)
		if err != nil {
			panic(err)
		}
		res := strconv.FormatInt(n1*i1, 10)
		for j := 0; j < zero; j++ {
			res = res + "0"
		}
		fmt.Printf("%v * %v = %v\n", n1, i1, res)
		zero++
		if i2, err := strconv.ParseInt(res, 10, 64); err == nil {
			ans += i2
		}
	}
	return strconv.FormatInt(ans, 10)
}

func main() {
	fmt.Println(
		multiply("498828660196", "840477629533"),
		//multiply("123", "456"),
		//multiply("3", "4"),
		//multiply("567", "478960"),
	)
	var a int64
	a = 419254329864656431168468
	fmt.Println(a)
}

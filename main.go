package main

import (
	"fmt"
)

func main() {
	//if len(os.Args) != 2 {
	//	os.Exit(1)
	//
	//}
	//fmt.Println("its over ", os.Args[1])
	//
	var power int
	power = 9000
	power1 := 8999
	power3 := getPower()

	fmt.Println("its over %d \n", power, power1, power3)

}

func getPower() int {
	return 9020

}

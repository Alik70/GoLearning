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
	name, power4 := "testname", 123987
	value, exists := set_power("goku")
	// in case which we dont need the first value we assign it to _
	//_, exists := set_power("goku")

	if exists {
		println("hey there", value)
	}

	fmt.Println("its over  \n", power, power1, power3, name, power4)

}

func getPower() int {
	return 9020

}

func log(message string) {

}

//or add(a, b int) if they have the same type
func add(a int, b int) int {
	return 1
}

func set_power(name string) (int, bool) {

}

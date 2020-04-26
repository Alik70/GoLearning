package main

import (
	"GoLearning/structs"
)

func main() {
	//if len(os.Args) != 2 {
	//	os.Exit(1)
	//
	//}
	//fmt.Println("its over ", os.Args[1])
	////
	//var power int
	//power = 9000
	//power1 := 8999
	//power3 := getPower()
	//name, power4 := "testname", 123987
	//value, exists := set_power("goku")
	//// in case which we dont need the first value we assign it to _
	////_, exists := set_power("goku")
	//
	//if exists {
	//	println("hey there", value)
	//}
	//
	//fmt.Println("its over  \n", power, power1, power3, name, power4)
	////////////////////////////////////////////////////////////////////////////// chapter 2
	// & gets the address of the struct built
	//goku := &structs.Saiyan{
	//	Name:  "GOku",
	//	Power: 9000,
	//}
	//
	//////goku1 := structs.Saiyan{}
	////goku1 := structs.Saiyan{Name: "test"}
	////goku1.Power = 123456
	//super_with_pointer(goku)
	//println(goku.Power)
	//hero := structs.Saiyan{"jondark", 9001}
	//hero.Super()
	//println(hero.Power)
	////	third_hero := new(structs.Saiyan) it allocates the memory required by a type
	//// third_hero.power = 123465
	gohan := &structs.Saiyan{
		Name:  "GOHAN",
		Power: 9001,
		Father: &structs.Saiyan{
			Name:   "Goku",
			Power:  9001213,
			Father: nil,
		},
	}
	println(gohan.Father.Power)

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
	return 1, false
}

//here the call was call by value not refrence so the changes was made to the copy of our object
func super(s structs.Saiyan) {
	s.Power += 1000

}

// & is address of operator which get the address of our value, and * means pointer to value fo type saiyan
func super_with_pointer(s *structs.Saiyan) {
	s.Power += 10000

}

// still we are givin a copy of the address to the func but it wont make a difference

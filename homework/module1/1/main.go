package main

import "fmt"

func main() {
	myTest := []string{"I", "am", "stupid", "and", "weak"}

	fmt.Printf("myTest %+v\n", myTest)
	for index, value := range myTest {
		if myTest[index] == "stupid" {
			fmt.Printf("modify %+v to %+v\n", value, "smart")
			myTest[index] = "smart"
		}
		if myTest[index] == "weak" {
			fmt.Printf("modify %+v to %+v\n", value, "strong")
			myTest[index] = "strong"
		}
	}
	fmt.Printf("myTest %+v\n", myTest)
}

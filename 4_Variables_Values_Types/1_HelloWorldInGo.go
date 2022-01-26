package main

import "fmt"

func main() {
	fmt.Println("Go Programming")
	Mars()
	fmt.Println("Additional Stuff")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	Neptune()

}
func Mars() {
	fmt.Println("Gravity on Mars ", 3.721, "m/s Square")
}
func Neptune() {
	fmt.Println("Gravity on Neptune ", 11.15, "m/s Square")
}

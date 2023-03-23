// Created By: Lamees Hemdan
// Created On: March 2023
//
// This program calcualtes area of a triangle

package main

import "fmt"

func main() {
	var base int
	var height int
	var area int
	// input
	fmt.Println("This program calculates the area of a triangle")
	fmt.Println("Enter the base (cm):")
	fmt.Scanln(&base)
	fmt.Println("Enter the height (cm):")
	fmt.Scanln(&height)
	// process
	area = (base * height) / 2
	// output
	fmt.Println("The area of the triangle is", area, " cm².")
	fmt.Println("\nDone.")
}

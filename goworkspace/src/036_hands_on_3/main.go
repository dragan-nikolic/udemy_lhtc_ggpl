/*
Use var to DECLARE three VARIABLES. The variables should have package level scope. 
Do not assign VALUES to the variables. 

Use the following IDENTIFIERS for the variables and make sure the variables are 
of the following TYPE (meaning they can store VALUES of that TYPE):

	identifier x type int
	identifier y type string
	identifier z type bool

In func main:
	- Print out the values for each identifier
	- The compiler assigned values to the variables. What are these values called?

At the package level scope, assign the following values to the three variables:
	for x assign 42
	for y assign "James Bond"
	for z assign true

In func main:
	- use fmt.Sprintf to print all of the VALUES to one single string. 
	  ASSIGN the returned value of TYPE string using the short declaration 
	  operator to a VARIABLE with the IDENTIFIER "s"
	- print out the value stored by variable "s"
*/

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)

	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println("s =", s)
}

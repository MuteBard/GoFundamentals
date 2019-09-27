//In the world of Go, there are two types of packages:

//- Executable (E)
//	- when compiled, spits out a runnable, executable file
//	- go build main.go
//	- used for actually doing something
//- Reusable (R)
//	- code dependencies or libraries
//	- we put in alot of reusable logic, helper functions, stuff that will help us in future projects

//How do we know we are making Executable or Reusable
//The name of the package you use determines whether you are making an executable or a dependency type package
//Specifically the word main is used to make a Executable type package

//package main = Defines a package that can be compiled and then "executed". Must have a func called 'main'
//package calculator = Defines a package that can be "reused" as a dependency (helper code)
//package uploader = Defines a package that can be "reused" as a dependency
package main

//fmt is the name of a standard library package that is included in go. fmt = format
//It is used to print out to the terminal
//other standard packages include main, math, encoding, debug, crypto, io
//golang.org/pkg
import "fmt"

//when using the main package, we must define a function named "main" which is ran automatically when the program runs
func main(){
	fmt.Println("Hi there!")
}
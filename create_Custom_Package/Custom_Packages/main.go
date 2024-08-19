package main

import (
	custom1 "goHasan/Custom_Packages/custom"
)

// --------------------------- Custom Packages ----------------------------------------

// go mod init <module-name> --->> Creates a module file
// go mod tidy  ---->> When you are importing any github package dependency

func main() {
	custom1.Val = 10
	custom1.PrintValue("Hello Golang !")
}

package utils

import "fmt"


// Imprime mensaje de éxito ID: 001
func PrintSuccess(args ...string) {
 
	fmt.Println("001 - ", args)
	
}

// Imprime mensaje de advertencia ID: 002
func PrintWarning(args ...string) {
	
	fmt.Println("002 - ", args)
	
}


// Imprime mensaje de ayuda ID: 003
func PrintHelp(args ...string) {

	for _, arg := range args {
		fmt.Println("003 - ", arg)
	}
	
}

// Imprime mensaje de error ID: 004
func PrintError(args ...string) {
	
	for _, arg := range args {
		fmt.Println("004 - ",arg)
	}
	
}

package utils

import "fmt"


// Mensaje de éxito ID: 001
func PrintSuccess(args ...string) {
	fmt.Println("001 - ", args)
}

// Mnsaje de advertencia ID: 002
func PrintWarning(args ...string) {
	fmt.Println("002 - ", args)
}

// Mensaje de ayuda ID: 003
func PrintHelp(args ...string) {
	for _, arg := range args {
		fmt.Println("003 - ", arg)
	}
}

// Mnsaje de error ID: 004
func PrintError(args ...string) {
	for _, arg := range args {
		fmt.Println("004 - ",arg)
	}
}

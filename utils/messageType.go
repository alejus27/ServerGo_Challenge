package utils

import "fmt"


// PrintSuccess prints success message
func PrintSuccess(args ...string) {
 
	fmt.Println("001 - ", args)
	
}

// PrintWarning prints a warning.
func PrintWarning(args ...string) {
	
	fmt.Println("002 - ", args)
	
}


// PrintHelp calls fmt. Println with help.
func PrintHelp(args ...string) {

	for _, arg := range args {
		fmt.Println("003 - ", arg)
	}
	
}

// PrintError prints an error message.
func PrintError(args ...string) {
	
	for _, arg := range args {
		fmt.Println("004 - ",arg)
	}
	
}

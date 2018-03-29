package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	 "github.com/goinaction/code/chapter2/sample/search"
	"fmt"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")


   //  fmt.Println(filepath.Dir(os.Args[0]))

	fmt.Println(os.Getwd())
}

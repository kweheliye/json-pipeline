package main

import (
	"fmt"
	"json-pipeline/cmd"
)

func main() {
	cmd.Execute() // All CLI commands handled in cmd package
	fmt.Println("Pipeline completed successfully!")
}

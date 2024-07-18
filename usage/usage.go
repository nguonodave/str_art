package usage

import "fmt"

// print_usage dipalays the usage for ascii-art.
func PrintUsage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
}

// print_output_usage dipalays the usage for ascii-art-output.
func PrintOutputUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

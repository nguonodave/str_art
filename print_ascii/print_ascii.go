package print_ascii

import "fmt"

// PrintAscii displays on the terminal, the art of the provided string.
func PrintAscii(char_art_map map[rune][]string, str_item string) {
	for i := 0; i < 8; i++ {
		for _, char := range str_item {
			fmt.Print(char_art_map[char][i])
		}
		fmt.Println()
	}
}

package write_output

import (
	"io"
	"log"
)

func WriteOutput(char_art_map map[rune][]string, str_item string, file io.Writer) {
	for i := 0; i < 8; i++ {
		for _, char := range str_item {
			out := char_art_map[char][i]
			if _, write_err := io.WriteString(file, out); write_err != nil {
				log.Fatalf("Error writing to file: %s", write_err)
			}
		}
		if _, write_err := io.WriteString(file, "\n"); write_err != nil {
			log.Fatalf("Error writing to file: %s", write_err)
		}
	}
}

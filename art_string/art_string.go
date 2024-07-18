package art_string

// WriteOutput stores the string art to a specified file.
func ArtString(char_art_map map[rune][]string, str_item string) string {
	var art_str string
	for i := 0; i < 8; i++ {
		for _, char := range str_item {
			out := char_art_map[char][i]
			art_str += out
		}
		art_str += "\n"
	}

	return art_str
}

package map_rune_art

func MapRuneArt(lines []string, char_art_map map[rune][]string, current_ascii_char int) {
	/*
		MAPPING EACH ART CHARACTER TO A RUNE
		--> 32 is the ascii value of space (the first ascii printable characyter)
		--> in the map (char_art_map), store an array of 8 lines (the lines are the string art representation)
		--> the first i++ goes to the next art line, immediately the first line has been appended
		--> the second i++ skips the empty line after each character art
	*/

	for i := 1; i < len(lines); {
		art_char := []string{}
		for j := 0; j < 8; j++ {
			art_char = append(art_char, lines[i])
			i++
		}
		i++
		char_art_map[rune(current_ascii_char)] = art_char
		current_ascii_char++
	}
}

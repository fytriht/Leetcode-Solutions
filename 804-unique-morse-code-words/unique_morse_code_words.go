package solution

var m = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	tr := map[string]bool{}
	for _, w := range words {
		s := ""
		for _, b := range w {
			s += m[b-'a']
		}
		tr[s] = true
	}
	return len(tr)
}

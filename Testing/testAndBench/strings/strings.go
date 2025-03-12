package strings

import (
	"sort"
	"strings"
)

// Add concatène deux chaînes de caractères
func Add(s1, s2 string) string {
	return s1 + s2
}

// IsInStr vérifie si toFind est dans str
func IsInStr(toFind, str string) bool {
	return strings.Index(str, toFind) >= 0
}

// SortWords trie un slice de mots en ordre alphabétique
func SortWords(words []string) {
	sort.Strings(words)
}

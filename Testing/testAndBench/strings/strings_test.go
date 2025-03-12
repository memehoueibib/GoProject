package strings

import (
	"testing"
)

// ----------------------
//       Test Add
// ----------------------
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want string
	}{
		{"Concat simple", "Hello", "World", "HelloWorld"},
		{"Vide à gauche", "", "World", "World"},
		{"Vide à droite", "Hello", "", "Hello"},
		{"Deux vides", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("Add(%q, %q) = %q; want %q", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

// ----------------------
//    Benchmark Add
// ----------------------
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add("Hello", "World")
	}
}

// ----------------------
//    Test IsInStr
// ----------------------
func TestIsInStr(t *testing.T) {
	tests := []struct {
		name   string
		toFind string
		str    string
		want   bool
	}{
		{"Contient 'hello'", "hello", "hello world", true},
		{"Ne contient pas 'bye'", "bye", "hello world", false},
		{"Empty toFind => found at 0", "", "hello", true},
		{"Empty str => false sauf si toFind vide", "hello", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsInStr(tt.toFind, tt.str)
			if got != tt.want {
				t.Errorf("IsInStr(%q, %q) = %v; want %v", tt.toFind, tt.str, got, tt.want)
			}
		})
	}
}

// ----------------------
// Benchmark IsInStr
// ----------------------
func BenchmarkIsInStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsInStr("world", "hello world")
	}
}

// ----------------------
//    Test SortWords
// ----------------------
func TestSortWords(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"Déjà trié", []string{"apple", "banana"}, []string{"apple", "banana"}},
		{"Ordre inverse", []string{"z", "y", "x"}, []string{"x", "y", "z"}},
		{"Mélangé", []string{"banana", "apple", "pear"}, []string{"apple", "banana", "pear"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]string, len(tt.input))
			copy(data, tt.input)
			SortWords(data)
			for i := range data {
				if data[i] != tt.want[i] {
					t.Errorf("SortWords(%v) = %v; want %v", tt.input, data, tt.want)
					break
				}
			}
		})
	}
}

// ----------------------
// Benchmark SortWords
// ----------------------
func BenchmarkSortWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []string{"banana", "apple", "pear", "orange"}
		SortWords(data)
	}
}

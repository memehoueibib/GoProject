package maths

import (
	"io"
	"os"
	"strings"
	"testing"
)

// ----------------------
//       Test Add
// ----------------------
func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		n1, n2  int
		want    int
	}{
		{"1 + 1 = 2", 1, 1, 2},
		{"-5 + 3 = -2", -5, 3, -2},
		{"0 + 0 = 0", 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.n1, tt.n2)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.n1, tt.n2, got, tt.want)
			}
		})
	}
}

// ----------------------
//    Benchmark Add
// ----------------------
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(i, i+1)
	}
}

// ----------------------
//       Test Fibo
// ----------------------
// Fibo() ne renvoie rien, il affiche 50 nombres. On teste que 50 lignes sont bien imprimées.
func TestFibo(t *testing.T) {
	// On redirige la sortie standard vers un buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Fibo()

	// On restaure la sortie standard
	w.Close()
	os.Stdout = oldStdout

	// On lit le contenu imprimé
	out, _ := io.ReadAll(r)
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) != 50 {
		t.Errorf("Fibo() devrait imprimer 50 lignes, en a imprimé %d", len(lines))
	}
}

// ----------------------
//    Benchmark Fibo
// ----------------------
func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibo()
	}
}

// ----------------------
//    Test Ordered
// ----------------------
func TestOrdered(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    bool
	}{
		{"Déjà trié", []int{1, 2, 3, 4}, true},
		{"Un seul élément", []int{42}, true},
		{"Pas trié", []int{2, 1}, false},
		{"Vides", []int{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ordered(tt.numbers)
			if got != tt.want {
				t.Errorf("Ordered(%v) = %v; want %v", tt.numbers, got, tt.want)
			}
		})
	}
}

// ----------------------
// Benchmark Ordered
// ----------------------
func BenchmarkOrdered(b *testing.B) {
	// On crée un slice déjà trié de 10_000 éléments
	slice := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		slice[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Ordered(slice)
	}
}

// ----------------------
//   Test AscendingSort
// ----------------------
func TestAscendingSort(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    []int
	}{
		{"Déjà trié", []int{1, 2, 3}, []int{1, 2, 3}},
		{"Inversé", []int{3, 2, 1}, []int{1, 2, 3}},
		{"Mélangé", []int{5, 2, 4, 1}, []int{1, 2, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// On copie l'input pour ne pas le modifier directement
			data := make([]int, len(tt.input))
			copy(data, tt.input)

			AscendingSort(data)
			// Vérification
			for i := range data {
				if data[i] != tt.want[i] {
					t.Errorf("AscendingSort(%v) = %v; want %v", tt.input, data, tt.want)
					break
				}
			}
		})
	}
}

// ----------------------
// Benchmark AscendingSort
// ----------------------
func BenchmarkAscendingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []int{5, 2, 4, 1, 10, 9, 8, 7, 6, 3, 0}
		AscendingSort(data)
	}
}

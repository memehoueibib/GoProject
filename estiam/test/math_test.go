package main

import "testing"

func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("add(2, 3) = %d; attendu 5", result)
    }
}
package main

import "testing"

func TestSquare(t *testing.T) {
    result := Square(5)
    if result != 25 {
        t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 25)
    }
}


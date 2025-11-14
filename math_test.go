package main

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Fatal("expected 2 + 3 = 5")
    }
    if Add(-1, 1) != 0 {
        t.Fatal("expected -1 + 1 = 0")
    }
}

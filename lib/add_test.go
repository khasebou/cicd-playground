package lib

import "testing"

// TestAdd
func TestAdd(t *testing.T) {
    a := -12
	b := 123
	exp := a + b
    if ans := Add(a, b); (ans != exp) {
        t.Fatalf("Added %d to %d. expected %d, got %d", a ,b, exp, ans)
    }
}

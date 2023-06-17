package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a, b, exp := 2, 3, 5

	res := add(a, b)
	if exp != res {
		t.Errorf("expected %d, got %d.", exp, res)
	}
}
package greeting

import (
	"testing"
)

func TestHello(t *testing.T) {
	expect := "Halo, Insan Dikti!"
	if resp := Hello(); resp != expect {
		t.Errorf("Hello() = %q, want %q", resp, expect)
	}
}

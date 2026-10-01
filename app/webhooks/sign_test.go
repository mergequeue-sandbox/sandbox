package webhooks

import "testing"

func TestSignIsStable(t *testing.T) {
	a, b := Sign([]byte("k"), []byte("{}")), Sign([]byte("k"), []byte("{}"))
	if a != b || len(a) != len("sha256=")+64 {
		t.Errorf("Sign() = %q, %q", a, b)
	}
}

package dragonvalid

import "testing"

func TestValidRule(t *testing.T) {
	v := Text("").Required().hasLetter()
	t.Log(v.err)
}

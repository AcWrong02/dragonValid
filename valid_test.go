package dragonvalid

import "testing"

func TestValidRule(t *testing.T) {
	v := Text("123").Required().hasLetter()
	err := v.Error()
	t.Log("err---", err)
}

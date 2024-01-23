package dragonvalid

import "testing"

func TestHasLower(t *testing.T) {
	errMsg := "必须含有小写字母"
	v1 := Text("123").Required().hasLower(errMsg)
	// err := v1.Error()
	if v1.err != nil {
		t.Fatal(errMsg)
	}

	v2 := Text("123a").Required().hasLower(errMsg)
	err2 := v2.Error()
	if err2 != nil {
		t.Error(errMsg)
	}
}

func TestHasUpper(t *testing.T) {
	errMsg := "必须含有大写字母"
	v1 := Text("123").Required().hasUpper(errMsg)
	err := v1.Error()
	if err != nil {
		t.Fatal(errMsg)
	}

	v2 := Text("123A").Required().hasUpper(errMsg)
	err2 := v2.Error()
	if err2 != nil {
		t.Error(errMsg)
	}
}

func TestHasNumber(t *testing.T) {
	errMsg := "必须含有数字"
	v1 := Text("123").Required().HasNumber(errMsg)
	err := v1.Error()
	if err != nil {
		t.Fatal(errMsg)
	}

	v2 := Text("A").Required().HasNumber(errMsg)
	err2 := v2.Error()
	if err2 != nil {
		t.Error(errMsg)
	}
}

func TestHasSymbol(t *testing.T) {
	errMsg := "必须含有字符"
	v1 := Text("123+-").Required().HasSymbol(errMsg)
	err := v1.Error()
	if err != nil {
		t.Fatal(errMsg)
	}

	v2 := Text("A").Required().HasSymbol(errMsg)
	err2 := v2.Error()
	if err2 != nil {
		t.Error(errMsg)
	}
}

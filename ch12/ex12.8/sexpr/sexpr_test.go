package sexpr_test

import (
	"bytes"
	"testing"

	"github.com/DaneSpiritGOD/ex12.8/sexpr"
)

type foo struct {
	A string
	B int
}

func TestDecode(t *testing.T) {
	fooV := foo{"hello", 1}
	content, err := sexpr.Marshal(fooV)
	if err != nil {
		t.Errorf("marshal error: %v", err)
		return
	}

	var fooV2 foo
	decoder := sexpr.NewDecoder(bytes.NewReader(content))
	err = decoder.Decode(&fooV2)
	if err != nil {
		t.Errorf("decode error: %v", err)
		return
	}

	if fooV2.A != fooV.A {
		t.Errorf("a not equal")
		return
	}

	if fooV2.B != fooV.B {
		t.Errorf("b not equal")
		return
	}
}

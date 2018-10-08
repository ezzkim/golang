package utils

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 2, 3
	if 5 == Add(a, b) {
		t.Log("test add success")
	} else {
		t.Error("test add fail")
	}
	fmt.Println("add test end")
}

func TestSub(t *testing.T) {
	a, b := 2, 3
	if -1 == Sub(a, b) {
		t.Log("test sub success")
	} else {
		t.Error("test sub fail")
	}
	fmt.Println("sub test end")
}

func TestMul(t *testing.T) {
	a, b := 2, 3
	if 6 == Mul(a, b) {
		t.Log("test mul success")
	} else {
		t.Error("test mul fail")
	}
	fmt.Println("mul test end")
}

func TestDiv(t *testing.T) {
	a, b := 12, 3
	if 4 == Div(a, b) {
		t.Log("test div success")
	} else {
		t.Error("test div fail")
	}
	fmt.Println("div test end")
}

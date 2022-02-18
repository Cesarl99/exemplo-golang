package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		test{data: []int{10, 10, 10}, answer: 30},
		test{[]int{3, 6, 2}, 11},
		test{[]int{-4, 4, 0, 22}, 22},
	}
	for _, v := range tests {
		z := Soma(v.data...)
		if z != v.answer {
			t.Error("esperado:", v.answer, "recebido", z)
		}
	}
}

/*
func TestSoma(t *testing.T) {
	x := soma(36, 65, 542, 98)
	total := 747
	if x != total {
		t.Error("esperado:", total, "recebido", x)
	}
}
func TestSoma1(t *testing.T) {
	x := soma(36, 65, 542, 98)
	total := 746
	if x != total {
		t.Error("esperado:", total, "recebido", x)
	}
}

func TestMult(t *testing.T)
y := mult(10)
resultado :=361
if y != resultado{
	t.Error("esperado",resultado,"recebido",y)
}
*/

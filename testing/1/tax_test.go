package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expect := 5.0
	got := CalculateTax(amount)
	if got != expect {
		t.Errorf("Expected %f but got %f", expect, got)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount float64
		want   float64
	}
	table := []calcTax{
		{
			500.0, 5.0,
		},
		{
			1000.0, 10.0,
		},
		{
			1500.0, 10.0,
		},
		{
			0, 0,
		},
	}
	for _, item := range table {
		got := CalculateTax(item.amount)
		if got != item.want {
			t.Errorf("Expected %f but got %f", item.want, got)
		}
	}
}

//comando para rodar o benchmark
//go test -bench=.
//comando para rodar apenas o benchmark
// /go test -bench=. -run=^#
//comando para rodar bench com X vezes por teste
//go test -bench=. -run=^# -count=10
//comando para rodar bench com alocação de memória
//go test -bench=. -run=^# -benchmem
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

//go test -fuzz=. -run=^#
//go test -fuzz=. -fuzztime=5s -run=^#
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1500.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		got := CalculateTax(amount)
		if amount <= 0 && got != 0 {
			t.Errorf("Received %f but expected 0", got)
		}
		if amount > 20000 && got != 20 {
			t.Errorf("Received %f but expected 20", got)
		}
	})
}

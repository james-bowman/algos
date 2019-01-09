package dot

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"testing"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat/sampleuv"
)

type params struct {
	x    []float64
	indx []int
	y    []float64
}

func buildVectors(nnz int, dim int, rnd *rand.Rand) params {
	x := make([]float64, nnz)
	indx := make([]int, nnz)
	y := make([]float64, dim)

	sampleuv.WithoutReplacement(indx, dim, rnd)
	sort.Ints(indx)

	for i := range x {
		x[i] = rnd.Float64()
	}

	for i := range y {
		y[i] = rnd.Float64()
	}

	return params{
		x:    x,
		indx: indx,
		y:    y,
	}
}

func TestDot(t *testing.T) {
	inputs := []struct {
		x    []float64
		indx []int
		y    []float64
		exp  float64
	}{
		{
			x:    []float64{},
			indx: []int{},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			exp:  0,
		},
		{
			x:    []float64{2},
			indx: []int{2},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			exp:  4,
		},
		{
			x:    []float64{2, 3},
			indx: []int{2, 3},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			exp:  13,
		},
		{
			x:    []float64{2, 3, 4},
			indx: []int{2, 3, 4},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			exp:  29,
		},
		{
			x:    []float64{2, 3, 4, 5},
			indx: []int{2, 3, 4, 5},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			exp:  54,
		},
		{
			x:    []float64{2, 3, 4, 5, 6},
			indx: []int{2, 3, 4, 5, 6},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			exp:  90,
		},
		{
			x:    []float64{2, 3, 4, 5, 6, 7},
			indx: []int{2, 3, 4, 5, 6, 7},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			exp:  139,
		},
	}

	for i, input := range inputs {
		for _, f := range funcs {
			dot := f(input.x, input.indx, input.y)
			if input.exp != dot {
				fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
				t.Errorf("Test %d: %s failed: Expected %f but received %f", i+1, fname, input.exp, dot)
			}
		}
	}
}
func BenchmarkDot(b *testing.B) {
	rnd := rand.New(rand.NewSource(0))

	inputs := []params{
		buildVectors(10, 100, rnd),
		buildVectors(100, 1000, rnd),
		buildVectors(1000, 10000, rnd),
		buildVectors(10000, 100000, rnd),
		buildVectors(100000, 1000000, rnd),
	}

	for _, input := range inputs {
		for _, f := range funcs {
			fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			b.Run(fmt.Sprintf("%d/%d %s", len(input.x), len(input.y), fname), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(input.x, input.indx, input.y)
				}
			})
		}
	}
}

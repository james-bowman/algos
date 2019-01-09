package dot

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat/sampleuv"
)

type params struct {
	x    []float64
	indx []int
	y    []float64
	incy int
}

func buildVectors(nnz int, dim int, inc int, rnd *rand.Rand) params {
	x := make([]float64, nnz)
	indx := make([]int, nnz)
	y := make([]float64, dim*inc)

	sampleuv.WithoutReplacement(indx, dim, rnd)
	//sort.Ints(indx)

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
		incy: inc,
	}
}

func TestDot(t *testing.T) {
	inputs := []struct {
		x    []float64
		indx []int
		y    []float64
		incy int
		exp  float64
	}{
		{
			x:    []float64{},
			indx: []int{},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			incy: 1,
			exp:  0,
		},
		{
			x:    []float64{2},
			indx: []int{2},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			incy: 1,
			exp:  4,
		},
		{
			x:    []float64{2, 3},
			indx: []int{2, 3},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			incy: 1,
			exp:  13,
		},
		{
			x:    []float64{2, 3, 4},
			indx: []int{2, 3, 4},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			incy: 1,
			exp:  29,
		},
		{
			x:    []float64{2, 3, 4, 5},
			indx: []int{2, 3, 4, 5},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			incy: 1,
			exp:  54,
		},
		{
			x:    []float64{2, 3, 4, 5, 6},
			indx: []int{2, 3, 4, 5, 6},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			incy: 1,
			exp:  90,
		},
		{
			x:    []float64{2, 3, 4, 5, 6, 7},
			indx: []int{2, 3, 4, 5, 6, 7},
			y:    []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			incy: 1,
			exp:  139,
		},
		{
			x:    []float64{2, 3, 4, 5, 6, 7},
			indx: []int{2, 3, 4, 5, 6, 7},
			y: []float64{
				0, 5,
				1, 5,
				2, 5,
				3, 5,
				4, 5,
				5, 5,
				6, 5,
				7, 5,
				8, 5,
				9, 5,
				10, 5,
			},
			incy: 2,
			exp:  139,
		},
	}

	for i, input := range inputs {
		for _, f := range funcs {
			dot := f(input.x, input.indx, input.y, input.incy)
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
		buildVectors(10, 100, 1, rnd),
		buildVectors(100, 1000, 1, rnd),
		buildVectors(1000, 10000, 1, rnd),
		buildVectors(10000, 100000, 1, rnd),
		buildVectors(100000, 1000000, 1, rnd),
		buildVectors(10, 100, 100, rnd),
		buildVectors(100, 1000, 1000, rnd),
		buildVectors(1000, 10000, 10000, rnd),
		buildVectors(4000, 10000, 10000, rnd),
		//		buildVectors(10000, 100000, 100000, rnd),
	}

	for _, input := range inputs {
		for _, f := range funcs {
			fname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			b.Run(fmt.Sprintf("%d/%d %s", len(input.x), len(input.y), fname), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(input.x, input.indx, input.y, input.incy)
				}
			})
		}
	}
}

package dot

// Dot is a function type for sparse dot product operations
// (r <- x^T * y) and calculates the dot product of
// sparse vector x and dense vector y.  indx is used as the index
// values to gather.
type Dot func(x []float64, indx []int, y []float64, incy int) float64

var funcs = []Dot{
	Dot1,
	Dot2,
	Dot3,
	Dot4a,
	Dot4b,
	Dot5b,
	Dot6a,
	Dot6b,
	Dot7a,
	Dot7b,
	//	Dot7c,
	Dot7d,
}

// Dot1 is a basic pure Go implementation of sparse dot product.
func Dot1(x []float64, indx []int, y []float64, incy int) (dot float64) {
	for i, index := range indx {
		dot += x[i] * y[index*incy]
	}
	return
}

// Dot2 is a basic, literal assembler implementation of sparse dot product.
func Dot2(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot3 is Dot2() with loop inversion optimisation (changing for
// loop to a do...while loop wrapped in an if statement).
func Dot3(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot4a is Dot3() with loop reversal optimisation (negative index
// incrementing up towards zero rather than up from zero, reducing loop overhead).
func Dot4a(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot4b is Dot3() with loop reversal optimisation (decrementing down towards
// zero rather incrementing up from zero, reducing loop overhead).
func Dot4b(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot5b is Dot4b() with loop unrolling (x4) optimisation.
func Dot5b(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot6a is Dot5a() with software pipelining optimisations.
func Dot6a(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot6b is Dot5b() with software pipelining optimisations.
func Dot6b(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot7a is Dot6a() with vectorisation (SIMD) optimisations.
func Dot7a(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot7b is Dot6b() with vectorisation (SIMD) optimisations.
func Dot7b(x []float64, indx []int, y []float64, incy int) (dot float64)

// Dot7c is Dot6b() with vectorisation (SIMD) optimisations but without loop reversal optimisation.
func Dot7c(x []float64, indx []int, y []float64) (dot float64)

// Dot7d is Dot6b() with vectorisation (SIMD) optimisations but without pipelining optimisation.
func Dot7d(x []float64, indx []int, y []float64, incy int) (dot float64)

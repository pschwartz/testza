package testza

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

// MockInputsFloats64Helper contains integer test sets.
// Use testza.Use.Mock.Inputs.Floats64.
type MockInputsFloats64Helper struct{}

func (h MockInputsFloats64Helper) Full() (floats []float64) {
	for i := 0; i < 50; i++ {
		floats = append(floats,
			h.GenerateRandomPositive(1, float64(i*1000))[0],
			h.GenerateRandomNegative(1, float64(i*1000*-1))[0],
		)
	}
	return
}

// GenerateRandomRange generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func (h MockInputsFloats64Helper) GenerateRandomRange(count int, min, max float64) (floats []float64) {
	for i := 0; i < count; i++ {
		floats = append(floats, min+rand.Float64()*(max-min))
	}

	return
}

// GenerateRandomPositive generates random positive integers with a maximum of max.
// If the maximum is 0, or below, the maximum will be set to math.MaxInt64.
func (h MockInputsFloats64Helper) GenerateRandomPositive(count int, max float64) (floats []float64) {
	if max <= 0 {
		max = math.MaxFloat64
	}

	floats = append(floats, h.GenerateRandomRange(count, 0, max)...)

	return
}

// GenerateRandomNegative generates random negative integers with a minimum of min.
// If the minimum is positive, it will be converted to a negative number.
// If it is set to 0, there is no limit.
func (h MockInputsFloats64Helper) GenerateRandomNegative(count int, min float64) (floats []float64) {
	if min > 0 {
		min *= -1
	} else if min == 0 {
		min = math.MaxFloat64 * -1
	}

	floats = append(floats, h.GenerateRandomRange(count, min, 0)...)

	return
}

// RunTests runs a test for every value in a testset.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hunderts of cases easily.
func (s MockInputsFloats64Helper) RunTests(t testRunner, testSet []float64, testFunc func(t *testing.T, index int, f float64)) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	test := internal.GetTest(t)
	if test == nil {
		t.Error(internal.ErrCanNotRunIfNotBuiltinTesting)
		return
	}

	for i, v := range testSet {
		test.Run(fmt.Sprint(v), func(t *testing.T) {
			t.Helper()

			testFunc(t, i, v)
		})
	}
}

// Modify returns a modified version of a test set.
func (h MockInputsFloats64Helper) Modify(inputSlice []float64, f func(index int, value float64) float64) (floats []float64) {
	for i, input := range inputSlice {
		floats = append(floats, f(i, input))
	}

	return
}

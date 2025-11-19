package calc

import (
	"math"
	"math/rand"
	"testing"
)

func TestSqrt(t *testing.T) {
	testsCount := 10
	for i := 0; i < testsCount; i++ {
		var num float64 = math.Round(float64(rand.Intn(100)))
		actual := Sqrt(num)
		except := math.Sqrt(num)

		if math.Round(actual) != math.Round(except) {
			t.Errorf("ответ Sqrt %.6f не совпадает с ожидаемым %.6f", actual, except)
		}
	}
}

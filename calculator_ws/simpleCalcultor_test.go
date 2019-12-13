package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
  result := Sum(int(3), int(4))
  expecting := 8
  if result != expecting {
		t.Errorf("expecting %d, got %d", expecting, result)
  }
}
func TestSqrt(t *testing.T) {
  result := Sqrt(9)
  expecting := float64(3)
  if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
  }
  result = Sqrt(-4)
  expecting = 2
  if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
  }
}
func TestFactorial(t *testing.T) {
	result := Factorial(4)
  expecting := uint64(24)
  if result != expecting {
		t.Errorf("expecting %d, got %d", expecting, result)
	}
}
func TestLog(t *testing.T) {
  result := Log(1)
  expecting := float64(0)
  if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
  }
}

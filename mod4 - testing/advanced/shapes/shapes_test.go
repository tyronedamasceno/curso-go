package shapes_test

import (
	"advanced-test/shapes"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		r := shapes.Rectangle{10, 12}
		expected := float64(120)
		result := r.Area()

		if expected != result {
			t.Errorf("Expected %f is different of result %f", expected, result)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		c := shapes.Circle{10}
		expected := float64(math.Pi * 100)
		result := c.Area()

		if expected != result {
			t.Errorf("Expected %f is different of result %f", expected, result)
		}
	})
}

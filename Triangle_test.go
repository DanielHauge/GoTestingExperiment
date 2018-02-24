package GoTestingExperiment

import (
	"math/rand"
	"testing"
)

func TestMakeValidTriangle(t *testing.T) {

	err, tri := MakeTriangle(1, 1, 1)
	t.Log(tri.PrintInfo())
	if err != nil {
		t.Error(err)
	}

}

func TestMakeValidTriangle2(t *testing.T) {

	err, tri := MakeTriangle(4, 4, 4)
	t.Log(tri.PrintInfo())
	if err != nil {
		t.Error(err)
	}
}

func TestMakeInvalidTriangle(t *testing.T) {

	err, tri := MakeTriangle(1, 2, 3)
	t.Log(tri.PrintInfo())
	if err == nil {
		t.Error(err)
	}

}

func TestCalculateFailArea(t *testing.T) {
	_, tri := MakeTriangle(1.5, 3.8, 2.5)
	t.Log(tri.PrintInfo())
	if tri.Area != 1.14 {
		t.Error()
	}

}

func TestCalculateTriangleType(t *testing.T) {

	_, tri := MakeTriangle(4, 2, 4)
	t.Log(tri.PrintInfo())

	if tri.TriangleType != 1 {
		t.Error()
	}

}

func TestCalculateAngles(t *testing.T) {
	_, tri := MakeTriangle(1.5, 3.8, 2.5)
	t.Log(tri.PrintInfo())
	if !contains(tri.Angles, 23.68) {
		t.Error()
	}
	if !contains(tri.Angles, 142.37) {
		t.Error()
	}
	if !contains(tri.Angles, 13.95) {
		t.Error()
	}
}

func BenchmarkMakeTriangle(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MakeTriangle(((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1))
	}

}

func BenchmarkCalculateAngles(b *testing.B) {

	sides := []float64{((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1)}
	t := Triangle{sides, []float64{0, 0, 0}, 0, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Angles = CalculateAngles(sides)
	}

}

func BenchmarkCalculateArea(b *testing.B) {

	sides := []float64{((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1)}
	t := Triangle{sides, []float64{0, 0, 0}, 0, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Area = CalculateArea(sides)
	}

}

func BenchmarkCalculateTriangleType(b *testing.B) {

	sides := []float64{((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1), ((rand.Float64() * 100) + 1)}
	t := Triangle{sides, []float64{0, 0, 0}, 0, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, t.TriangleType = CalculateTriangleType(sides)
	}

}

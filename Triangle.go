package GoTestingExperiment

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Triangle struct {
	Sides        []float64
	Angles       []float64
	Area         float64
	TriangleType int
}

func MakeTriangle(a float64, b float64, c float64) (err error, t Triangle) {
	err = nil
	sides := []float64{a, b, c}
	angles := CalculateAngles(sides)
	area := CalculateArea(sides)
	err, ttype := CalculateTriangleType(sides)

	t = Triangle{sides, angles, area, ttype}

	return err, t
}

func CalculateTriangleType(s []float64) (err error, tt int) {
	// 0 Scalene
	// 1 Isolecas
	// 2 Equilateral

	err = nil
	if !(s[0]+s[1] > s[2] && s[0]+s[2] > s[1] && s[1]+s[2] > s[0] && s[0] > 0 && s[1] > 0 && s[2] > 0) {
		err = errors.New("Sides cannot make up a triangle")
	}

	if s[0] == s[1] || s[1] == s[2] || s[0] == s[2] {
		tt = 1
	} else if s[0] == s[1] && s[1] == s[2] {
		tt = 2
	} else {
		tt = 0
	}
	return err, tt
}

func CalculateAngles(sides []float64) []float64 {
	result := []float64{0, 0, 0}

	upper := math.Pow(sides[1], 2) + math.Pow(sides[2], 2) - math.Pow(sides[0], 2)
	lower := 2 * sides[1] * sides[2]
	radian := math.Acos(upper / lower)

	res, err := strconv.ParseFloat(fmt.Sprintf("%.2f", radian*(180.0/math.Pi)), 64)
	if err != nil {
		panic(err)
	}
	result[0] = res

	upper = math.Pow(sides[0], 2) + math.Pow(sides[1], 2) - math.Pow(sides[2], 2)
	lower = 2 * sides[0] * sides[1]
	radian = math.Acos(upper / lower)

	res, err = strconv.ParseFloat(fmt.Sprintf("%.2f", radian*(180.0/math.Pi)), 64)
	if err != nil {
		panic(err)
	}
	result[1] = res

	result[2] = 180 - result[0] - result[1]

	return result

}

func CalculateArea(sides []float64) float64 {

	temp := (sides[0] + sides[1] + sides[2]) / 2
	res, err := strconv.ParseFloat(fmt.Sprintf("%.2f", math.Sqrt(temp*(temp-sides[0])*(temp-sides[1])*(temp-sides[2]))), 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (t *Triangle) PrintInfo() string {

	var buffer bytes.Buffer

	buffer.WriteString("TriangleType: " + strconv.Itoa(t.TriangleType) + "\n")

	buffer.WriteString("Sides: | ")

	for _, i := range t.Sides {
		buffer.WriteString(strconv.FormatFloat(i, 'f', -1, 64))
		buffer.WriteString(" | ")
	}
	buffer.WriteString("\n")
	buffer.WriteString("Angles: | ")
	for _, i := range t.Angles {
		buffer.WriteString(strconv.FormatFloat(i, 'f', -1, 64))
		buffer.WriteString(" | ")
	}
	buffer.WriteString("\n")
	buffer.WriteString("\nArea ")
	buffer.WriteString(strconv.FormatFloat(t.Area, 'f', -1, 64))
	return buffer.String()

}

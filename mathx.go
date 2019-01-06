// Copyright 2019 Wilhelm Stoll. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Pure go implementation of additional math functions.

package mathx

import (
	"math"
)

// Deg converts radians to degree.
func Deg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

// Rad converts degree to radians.
func Rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// Add calculates the add product for two vectors.
func Add(x, y [3]float64) [3]float64 {
	return [3]float64{
		x[0] + y[0],
		x[1] + y[1],
		x[2] + y[2],
	}
}

// Subtract calculates the subtract product for two vectors.
func Subtract(x, y [3]float64) [3]float64 {
	return [3]float64{
		x[0] - y[0],
		x[1] - y[1],
		x[2] - y[2],
	}
}

// Multiply calculates the multiply product for two vectors.
func Multiply(x [3]float64, y float64) [3]float64 {
	return [3]float64{
		x[0] * y,
		x[1] * y,
		x[2] * y,
	}
}

// Divide calculates the divide product for two vectors.
func Divide(x [3]float64, y float64) [3]float64 {
	return [3]float64{
		x[0] / y,
		x[1] / y,
		x[2] / y,
	}
}

// Cross calculates the cross product for two vectors.
func Cross(x, y [3]float64) [3]float64 {
	return [3]float64{
		x[1]*y[2] - x[2]*y[1],
		x[2]*y[0] - x[0]*y[2],
		x[0]*y[1] - x[1]*y[0],
	}
}

// Dot calculates the dot product of two vectors.
func Dot(x, y [3]float64) float64 {
	return x[0]*y[0] + x[1]*y[1] + x[2]*y[2]
}

// Norm calculates the norm of the number.
func Norm(x [3]float64) float64 {
	return math.Sqrt(Square(x[0]) + Square(x[1]) + Square(x[2]))
}

// Square multiplies the number by itself.
func Square(x float64) float64 {
	return x * x
}

// Sign indicates whether the number specified is negative or positive.
func Sign(x float64) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	}

	return 0
}

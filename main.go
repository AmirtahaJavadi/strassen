package main

import (
	"fmt"
	"time"
)

func add(a, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

// Subtract two matrices
func subtract(a, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = a[i][j] - b[i][j]
		}
	}
	return result
}
func MakeRow(s int) []int {
	var row []int
	for k := 1; k < s+1; k++ {
		row = append(row, k)
	}
	return row
}
func MakeMatrix(s int) [][]int {
	var matrix [][]int
	row := MakeRow(s)
	for k := 0; k < s; k++ {
		matrix = append(matrix, row)
	}
	return matrix
}
func standardMultiply(a, b [][]int) [][]int {
	n := len(a)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func strassenMultiply(a, b [][]int) [][]int {
	n := len(a)

	// Base case: if matrix is small, use standard multiplication
	if n <= 2 {
		return standardMultiply(a, b)
	}

	// Split matrices into quarters
	half := n / 2

	a11 := make([][]int, half)
	a12 := make([][]int, half)
	a21 := make([][]int, half)
	a22 := make([][]int, half)

	b11 := make([][]int, half)
	b12 := make([][]int, half)
	b21 := make([][]int, half)
	b22 := make([][]int, half)

	for i := 0; i < half; i++ {
		a11[i] = a[i][:half]
		a12[i] = a[i][half:]
		a21[i] = a[i+half][:half]
		a22[i] = a[i+half][half:]

		b11[i] = b[i][:half]
		b12[i] = b[i][half:]
		b21[i] = b[i+half][:half]
		b22[i] = b[i+half][half:]
	}

	// Calculate the 7 products recursively
	p1 := strassenMultiply(a11, subtract(b12, b22))
	p2 := strassenMultiply(add(a11, a12), b22)
	p3 := strassenMultiply(add(a21, a22), b11)
	p4 := strassenMultiply(a22, subtract(b21, b11))
	p5 := strassenMultiply(add(a11, a22), add(b11, b22))
	p6 := strassenMultiply(subtract(a12, a22), add(b21, b22))
	p7 := strassenMultiply(subtract(a11, a21), add(b11, b12))

	// Calculate the quadrants of the result matrix
	c11 := add(subtract(add(p5, p4), p6), p2)
	c12 := add(p1, p2)
	c21 := add(p3, p4)
	c22 := subtract(subtract(add(p5, p1), p3), p7)

	// Combine the quadrants into the result matrix
	result := make([][]int, n)
	for i := 0; i < half; i++ {
		result[i] = append(c11[i], c12[i]...)
		result[i+half] = append(c21[i], c22[i]...)
	}

	return result
}
func main() {

	for j := 1; j <= 25; j++ {
		a := MakeMatrix(j)
		b := MakeMatrix(j)

		start := time.Now()
		result := standardMultiply(a, b)
		for i := range result {
			fmt.Println(result[i])
		}
		fmt.Println(time.Since(start))
		start2 := time.Now()
		result = strassenMultiply(a, b)
		for i := range result {
			fmt.Println(result[i])
		}
		fmt.Println(time.Since(start2))
	}

}

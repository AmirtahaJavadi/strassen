package matrix

import "fmt"

type Matrix struct {
	data [][]int
}

func MakeRow(s int) []int {
	var row []int
	for k := 1; k < s+1; k++ {
		row = append(row, k)
	}
	return row
}

func CreateMatrix(size int) Matrix {
	var matrix [][]int
	row := MakeRow(size)
	for k := 0; k < size; k++ {
		matrix = append(matrix, row)
	}
	return Matrix{data: matrix}
}

func add(a, b Matrix) Matrix {
	n := len(a.data)
	result := Matrix{}
	for i := 0; i < n; i++ {
		result.data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result.data[i][j] = a.data[i][j] + b.data[i][j]
		}
	}
	return result
}

// Subtract two matrices
func subtract(a, b Matrix) Matrix {
	n := len(a.data)
	result := Matrix{}
	for i := 0; i < n; i++ {
		result.data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result.data[i][j] = a.data[i][j] - b.data[i][j]
		}
	}
	return result
}

func (m *Matrix) PrintMatrix() {
	for i := range m.data {
		fmt.Println(m.data[i])
	}
}

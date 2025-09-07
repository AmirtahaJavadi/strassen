package matrix

import "fmt"

type Matrix struct {
	data [][]int
}

func MakeRow(s int) []int {
	row := make([]int, s)
	for k := 0; k < s; k++ {
		row[k] = k + 1
	}
	return row
}

func CreateMatrix(size int) Matrix {
	matrix := make([][]int, size)
	row := MakeRow(size)
	for k := 0; k < size; k++ {
		newRow := make([]int, size)
		copy(newRow, row)
		matrix[k] = newRow
	}
	return Matrix{data: matrix}
}

func Add(a, b Matrix) Matrix {
	n := len(a.data)
	result := Matrix{data: make([][]int, n)}
	for i := 0; i < n; i++ {
		result.data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result.data[i][j] = a.data[i][j] + b.data[i][j]
		}
	}
	return result
}

func Subtract(a, b Matrix) Matrix {
	n := len(a.data)
	result := Matrix{data: make([][]int, n)}
	for i := 0; i < n; i++ {
		result.data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result.data[i][j] = a.data[i][j] - b.data[i][j]
		}
	}
	return result
}

func (m *Matrix) Data() [][]int { return m.data }

func (m *Matrix) PrintMatrix() {
	for i := range m.data {
		fmt.Println(m.data[i])
	}
}

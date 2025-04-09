package main

import (
	"fmt"
	"time"
)

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
func Multiplie(matrix1, matrix2 [][]int) [][]int {
	return nil
}

func StrassenMulti(matrix1, matrix2 [][]int) [][]int {
	return nil
}
func main() {

	for j := 1; j <= 25; j++ {
		a := MakeMatrix(j)
		b := MakeMatrix(j)
		size := len(a)
		result := make([][]int, size)
		start := time.Now()
		for i := range result {
			result[i] = make([]int, size)
		}

		for i := 0; i < size; i++ {
			for k := 0; k < size; k++ {
				for j := 0; j < size; j++ {
					result[i][j] += a[i][k] * b[k][j]
				}
			}
		}
		for i := range result {
			fmt.Println(result[i])
		}
		fmt.Println(time.Since(start))
	}

}

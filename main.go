package main

import "fmt"

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
func main() {
	matrix := [][]int{
		{1, 2, 3}, {1, 2, 3}, {1, 2, 3},
	}
	for j := 1; j <= 10; j++ {
		matrix = MakeMatrix(j)
		for i := range matrix {
			fmt.Println(matrix[i])
		}
	}
}

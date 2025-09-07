package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func makeRow(s int) []int {
	row := make([]int, s)
	for k := 0; k < s; k++ {
		row[k] = k + 1
	}
	return row
}

func makeMatrix(s int) [][]int {
	m := make([][]int, s)
	r := makeRow(s)
	for i := 0; i < s; i++ {
		row := make([]int, s)
		copy(row, r)
		m[i] = row
	}
	return m
}

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

// --- Strassen (with padding) ---

func nextPow2(n int) int {
	if n <= 1 {
		return 1
	}
	p := 1
	for p < n {
		p <<= 1
	}
	return p
}

func padMatrix(m [][]int, size int) [][]int {
	n := len(m)
	if n == size {
		cp := make([][]int, n)
		for i := range m {
			cp[i] = append([]int(nil), m[i]...)
		}
		return cp
	}
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
	}
	for i := 0; i < n; i++ {
		copy(res[i], m[i])
	}
	return res
}

func unpadMatrix(m [][]int, n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = append([]int(nil), m[i][:n]...)
	}
	return res
}

func strassenCore(a, b [][]int, threshold int) [][]int {
	n := len(a)
	if n <= threshold {
		return standardMultiply(a, b)
	}
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

	p1 := strassenCore(a11, subtract(b12, b22), threshold)
	p2 := strassenCore(add(a11, a12), b22, threshold)
	p3 := strassenCore(add(a21, a22), b11, threshold)
	p4 := strassenCore(a22, subtract(b21, b11), threshold)
	p5 := strassenCore(add(a11, a22), add(b11, b22), threshold)
	p6 := strassenCore(subtract(a12, a22), add(b21, b22), threshold)
	p7 := strassenCore(subtract(a11, a21), add(b11, b12), threshold)

	c11 := add(subtract(add(p5, p4), p6), p2)
	c12 := add(p1, p2)
	c21 := add(p3, p4)
	c22 := subtract(subtract(add(p5, p1), p3), p7)

	result := make([][]int, n)
	for i := 0; i < half; i++ {
		rowTop := append(append([]int(nil), c11[i]...), c12[i]...)
		rowBot := append(append([]int(nil), c21[i]...), c22[i]...)
		result[i] = rowTop
		result[i+half] = rowBot
	}
	return result
}

func strassenMultiply(a, b [][]int, threshold int) [][]int {
	n := len(a)
	m := nextPow2(n)
	ap := padMatrix(a, m)
	bp := padMatrix(b, m)
	cp := strassenCore(ap, bp, threshold)
	return unpadMatrix(cp, n)
}

func main() {
	// Choose sizes to test
	sizes := []int{2, 4, 8, 16, 32, 64, 96, 128, 192, 256}

	// CSV file to save timings
	f, err := os.Create("timings.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	w.Write([]string{"n", "standard_ms", "strassen_ms"})

	for _, n := range sizes {
		a := makeMatrix(n)
		b := makeMatrix(n)

		t0 := time.Now()
		_ = standardMultiply(a, b)
		t1 := time.Now()
		_ = strassenMultiply(a, b, 64) // threshold can be tuned
		t2 := time.Now()

		stdMs := float64(t1.Sub(t0).Microseconds()) / 1000.0
		strMs := float64(t2.Sub(t1).Microseconds()) / 1000.0

		fmt.Printf("n=%d  standard=%0.3f ms  strassen=%0.3f ms\n", n, stdMs, strMs)
		w.Write([]string{
			fmt.Sprint(n),
			fmt.Sprintf("%.3f", stdMs),
			fmt.Sprintf("%.3f", strMs),
		})
	}
}

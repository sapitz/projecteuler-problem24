package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const dim int = 10

var a = [dim]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var pos int = 3000000

var perm, facts [dim]int

func main() {

	for i := 0; i < dim; i++ {
		facts[i] = int(factorial(i))
	}

	fmt.Printf("This is the original set: \n")
	fmt.Printf("[%s] \n", iarrtostr(a))

	fmt.Printf("Factorials for dim = %d: \n", dim)
	fmt.Printf("[%s] \n", iarrtostr(facts))

	// fmt.Printf("Permutations generated with Heap algorithm: \n")
	// perm = a
	// permuteHeap(dim, &perm)

	// fmt.Printf("Permutations generated with Recursive Backtrack algorithm: \n")
	// perm = a
	// permute(0, &perm)

	//fmt.Printf("Permutations generated in Lexicographic sequence: \n")
	//perm = a
	//permuteLexi(&perm)

	fmt.Printf("Permutation at position in Lexicographic sequence: \n")
	perm = a
	getLexiPerm(perm, pos)

}

func getLexiPerm(arr [dim]int, pos int) {
	mod := pos
	res := arr
	for i := dim - 1; i >= 0; i-- {
		res[i] = int(math.Trunc(float64(mod / facts[i])))
		mod = mod % facts[i]
		fmt.Printf("%d: %d \n", i, mod)
		//	if mod == 0 {
		//		break
		//	}
	}
	reverse(&res, 0, dim-1)
	fmt.Printf("[%s] \n", iarrtostr(res))
	fmt.Printf("[%s] \n", iarrtostr(arr))
	for i := 0; i < dim; i++ {
		temp := res[i]
		res[i] = arr[res[i]]
		remove(&arr, temp)
		fmt.Println()
		fmt.Printf("%d: [%s] \n", i, iarrtostr(res))
		fmt.Printf("[%s] \n", iarrtostr(arr))
	}
}

func permuteLexi(arr *[dim]int) {
	p := 0
	for n := 1; n <= int(factorial(dim)); n++ {
		if n == pos {
			fmt.Printf("%d: [%s] \n", n, iarrtostr(*arr))
		}
		for i := dim - 2; i >= 0; i-- {
			if arr[i] < arr[i+1] {
				p = i
				break
			}
		}
		for j := dim - 1; j >= 0; j-- {
			if arr[j] > arr[p] {
				swap(arr, p, j)
				reverse(arr, p+1, dim-1)
				break
			}
		}
	}
}

func permute(p int, arr *[dim]int) {
	if p == dim-1 {
		fmt.Printf("[%s] \n", iarrtostr(*arr))
	} else {
		for i := p; i < dim; i++ {
			swap(arr, p, i)
			permute(p+1, arr)
			swap(arr, p, i)
		}
	}
}

func permuteHeap(p int, arr *[dim]int) {
	if p == 1 {
		fmt.Printf("[%s] \n", iarrtostr(*arr))
	} else {
		permuteHeap(p-1, arr)
		for i := 0; i < p-1; i++ {
			if p%2 == 0 {
				swap(arr, i, p-1)
			} else {
				swap(arr, 0, p-1)
			}
			permuteHeap(p-1, arr)
		}
	}
}

func swap(a *[dim]int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func reverse(a *[dim]int, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		swap(a, i, j)
	}
}

func remove(a *[dim]int, pos int) {
	for i := pos; i < dim-1; i++ {
		a[i] = a[i+1]
	}
	a[dim-1] = 0
}

/*
function to calculate n!
	n is an integer
	returns an unsigned 64 bit integer value to support large results
	if n is negative, returns 0
*/
func factorial(n int) uint64 {
	var fact uint64 = 1

	if n < 0 {
		fact = 0
	} else {
		for i := 1; i <= n; i++ {
			fact *= uint64(i)
		}
	}
	return fact
}

/*
function to convert an array of integers to a string
	a is an array of integers of size dim
	returns a string with space delimited array elements
*/
func iarrtostr(a [dim]int) string {
	str := make([]string, len(a))
	for i, v := range a {
		str[i] = strconv.Itoa(v)
	}
	return strings.Join(str, " ")
}

/*
https://www.bernardosulzbach.com/lexicographic-permutations/
http://www.tropicalcoder.com/APermutationOnCombinatorialAlgorithms.htm
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const dim int = 10

var a = [dim]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var pos int = 1000000

var perm, facts [dim]int

func main() {

	for i := 0; i < dim; i++ {
		facts[i] = int(factorial(i))
	}

	// fmt.Printf("This is the original set: \n")
	// fmt.Printf("[%s] \n", iarrtostr(a))

	// fmt.Printf("Factorials for dim = %d: \n", dim)
	// fmt.Printf("[%s] \n", iarrtostr(facts))

	// fmt.Printf("Permutations generated with Heap algorithm: \n")
	// perm = a
	// permuteHeap(dim, &perm)

	// fmt.Printf("Permutations generated with Recursive Backtrack algorithm: \n")
	// perm = a
	// permute(0, &perm)

	fmt.Printf("Permutations generated in Lexicographic sequence: \n")
	perm = a
	permuteLexi(&perm)

	fmt.Printf("Permutation at position %d in Lexicographic sequence: \n", pos)
	perm = a
	getLexiPerm(perm, pos)

}

/*
function to directly calculate and print the permutation of elements at a given position of
all possible, lexicographically ordered permutations of a given array of elements, as per problem 24
of the Euler Project.
	arr is an arr of integers of size dim
	pos is the position of the permutation to calculate

The alogorithm developed and used here is illustrated by example in the accompanying spreadsheet.
There are 2 steps to this approach:
	1.	Use size of arr and calculate whole integer of pos divided by factorial of each size-value -1.
		Store result into a result array. Note: same index is used here, so results are backwards.
		Iterate until result is 0, using remainder (mod) of each calculation as input into the next one.
		Use position where remainder (mod) result is 0 as reversal point. Stop and reverse results array.
	2. 	Iterate over results array.
		Use each result (trunc) value as an index into position array. Remove any positions allocated.
		At reversal point value (index) is one less.
		Reverse remainder of results array and complete iterations.
*/
func getLexiPerm(arr [dim]int, pos int) {
	var res [dim]int
	var mod0, temp int

	for i := dim - 1; i >= 0; i-- {
		res[i] = int(math.Trunc(float64(pos / facts[i])))
		pos = pos % facts[i]
		if pos == 0 {
			mod0 = dim - 1 - i
			break
		}
	}
	reverse(&res, 0, dim-1)
	for i := 0; i < dim; i++ {
		if i == mod0 {
			temp = res[i] - 1
		} else {
			temp = res[i]
		}
		res[i] = arr[temp]
		remove(&arr, temp)
		if i == mod0 {
			reverse(&arr, 0, dim-2-i)
		}
		// Uncomment next line to see how solution develops per iteration
		fmt.Printf("%d: [%s] [%s] \n", i, iarrtostr(res), iarrtostr(arr))
	}
	fmt.Printf("[%s] \n", iarrtostr(res))
}

/*
function to develop all permutations of an array of integers in lexicographical order and print value at pos.
This algorithm is used to validate direct approach, taken from https://www.bernardosulzbach.com/lexicographic-permutations/
*/
func permuteLexi(arr *[dim]int) {
	p := 0
	for n := 1; n <= int(factorial(dim)); n++ {
		// Print only value at pos
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

/*
function to swap two elements in an array
		a is a reference to an array of size dim
		i, j specifiy 0-based position of elements to swap
*/
func swap(a *[dim]int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

/*
function to reverse order of a contiguous portion of array
		a is a reference to an array of size dim
		start, end specifiy 0-based sub array to reverse
*/
func reverse(a *[dim]int, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		swap(a, i, j)
	}
}

/*
function to remove item at specified position in an array and pad 0 at end of array
	a is a reference to an array of size dim
	pos position of element to remove
*/
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
http://www.tropicalcoder.com/APermutationOnCombinatorialAlgorithms.htm
*/

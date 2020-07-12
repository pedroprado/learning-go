package sort

import (
	"reflect"
	"testing"
)

//-------------------------BENCHMARK---------------------------
//Benchmark : testing the performance of a function
func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{9, 6, 2, 4}

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}

}

func BenchmarkNormalSort(b *testing.B) {
	arr := []int{9, 6, 2, 4}

	for i := 0; i < b.N; i++ {
		NormalSort(arr)
	}
}

func BenchmarkBubbleSort_2(b *testing.B) {
	arr := makeArray(10000)

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}

}

func BenchmarkNormalSort_2(b *testing.B) {
	arr := makeArray(10000)

	for i := 0; i < b.N; i++ {
		NormalSort(arr)
	}
}

func BenchmarkBubbleSort_3(b *testing.B) {
	arr := makeArray(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}

}

//After 10 thousand elements, the built in Go function SORT is better than buble sort!
func BenchmarkNormalSort_3(b *testing.B) {
	arr := makeArray(100000)

	for i := 0; i < b.N; i++ {
		NormalSort(arr)
	}
}

//-----------------------------------------------
func TestMakeArray(t *testing.T) {
	arr := makeArray(10)

	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	equals := reflect.DeepEqual(arr, expected)

	if !equals {
		t.Errorf("\nExpected: %+v \n Got: %+v", expected, arr)
	}
}

//Creates an array of integers of size n
func makeArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i - 1
	}
	return arr

}

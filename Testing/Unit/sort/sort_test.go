package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 6, 2, 4}

	BubbleSort(arr)

	expected := []int{2, 4, 6, 9}

	equals := reflect.DeepEqual(arr, expected)

	if !equals {
		t.Errorf("\nExpected: %+v \n Got: %+v", expected, arr)
	}

}

func TestBubbleSort_AlreadySorted(t *testing.T) {
	arr := []int{2, 4, 6, 9}

	sorted := BubbleSort(arr)

	if !sorted {
		t.Errorf("\nShould be sorted.")
	}
}

func TestNormalSort(t *testing.T) {
	arr := []int{2, 8, 6, 5}

	NormalSort(arr)

	expected := []int{2, 5, 6, 8}

	equals := reflect.DeepEqual(arr, expected)

	if !equals {
		t.Errorf("\nExpected: %+v \n Got: %+v", expected, arr)
	}
}

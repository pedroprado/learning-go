package service

import (
	"reflect"
	"testing"
)

//This test case covers the integration between: service and sort layers
//The integration Test makes the ACTUAL CALL to the other layers (they should not be mocked)
func TestSort(t *testing.T) {
	arr := []int{6, 2, 4, 3}

	Sort(arr)

	expected := []int{2, 3, 4, 6}
	equals := reflect.DeepEqual(arr, expected)

	if !equals {
		t.Errorf("\nExpected: %+v \n Got: %+v \n", expected, arr)
	}
}

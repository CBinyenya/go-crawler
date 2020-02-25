package main

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	s := "One two three one two three one one"
	expected := map[string]int{
		"One":   1,
		"two":   2,
		"three": 2,
		"one":   3,
	}
	var actual = wordCounter(s)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expected, actual)
	}
}

func TestWordSorter(t *testing.T) {
	s := "One two three one two three one one"
	expectedSlice := []int{3, 2, 1}
	var words = wordCounter(s)
	actualMap, actualSlice := sorter(words)
	var expectedMap = map[int][]string{
		1: []string{"One"},
		2: []string{"two", "three"},
		3: []string{"one"},
	}
	if !reflect.DeepEqual(actualMap, expectedMap) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expectedMap, actualMap)
	}
	if !reflect.DeepEqual(actualSlice, expectedSlice) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expectedSlice, actualSlice)
	}

}

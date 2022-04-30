package main

import "testing"

func TestSum(t *testing.T) {
	cases := []struct{ x, y, expected int }{
		{2, 2, 4},
		{15, 2, 17},
		{2, 7, 9},
	}

	for _, testcase := range cases {
		actual := Sum(testcase.x, testcase.y)
		if actual != testcase.expected {
			t.Errorf("Sum was incorrect, expected %d but got %d", testcase.expected, actual)
		}
	}
}

func TestGetMax(t *testing.T) {
	cases := []struct {
		x, y, expected int
	}{
		{2, 2, 2},
		{15, 2, 15},
		{2, 7, 7},
	}

	for _, testcase := range cases {
		actual := GetMax(testcase.x, testcase.y)
		if actual != testcase.expected {
			t.Errorf("GetMax was incorrect, expected %d but got %d", testcase.expected, actual)
		}
	}
}

var fibonacciCases = []struct {
	n, expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{50, 12586269025},
}

func TestFibonacci(t *testing.T) {
	for _, testcase := range fibonacciCases {
		actual := Fibonacci(testcase.n)
		if actual != testcase.expected {
			t.Errorf("Fibonacci(%d) was incorrect, expected %d but got %d", testcase.n, testcase.expected, actual)
		}
	}
}

func TestComputeOnceFibonacci(t *testing.T) {
	for _, testcase := range fibonacciCases {
		actual := ComputeOnceFibonacci(testcase.n)
		if actual != testcase.expected {
			t.Errorf("Fibonacci(%d) was incorrect, expected %d but got %d", testcase.n, testcase.expected, actual)
		}
	}
}

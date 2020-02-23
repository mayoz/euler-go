package main

import "testing"

type TestCase struct {
    limit, out int
}

func TestSumEvenFibonacciNumbers(t *testing.T) {
    testPairs := []TestCase{
        {1, 0},
        {2, 2},
        {3, 2},
        {100, 44},
        {4000000, 4613732},
    }

    for i, v := range testPairs {
        actual := sumEvenFibonacciNumbers(v.limit)

        if actual != v.out {
            t.Errorf("#%d: sumEvenFibonacciNumbers(%d)=%d; expected: %d", i, v.limit, actual, v.out)
        }
    }
}

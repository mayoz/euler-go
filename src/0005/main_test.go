package main

import "testing"

type TestCase struct {
    goal   int
    output int
}

func TestSmallestNumber(t *testing.T) {
    testPairs := []TestCase{
        {10, 2520},
        {20, 232792560},
    }

    for i, v := range testPairs {
        actual := SmallestEvenlyDivisible(v.goal)

        if actual != v.output {
            t.Errorf("#%d: SmallestNumber(%d)=%d; expected: %d", i, v.goal, actual, v.output)
        }
    }
}

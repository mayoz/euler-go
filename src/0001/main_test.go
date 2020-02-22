package main

import "testing"

type TestCase struct {
    limit, out int
}

func TestSumMultiplies3or5(t *testing.T) {
    testPairs := []TestCase{
        {10, 23},
        {1000, 233168},
    }

    for i, v := range testPairs {
        actual := sumMultiplies3or5(v.limit)

        if actual != v.out {
            t.Errorf("#%d: sumMultiplies3or5(%d)=%d; expected: %d", i, v.limit, actual, v.out)
        }
    }
}

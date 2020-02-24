package main

import "testing"

type TestCase struct {
    input, output int64
}

func TestLargestPrimeFactor(t *testing.T) {
    testPairs := []TestCase{
        {13195, 29},
        {600851475143, 6857},
    }

    for i, v := range testPairs {
        actual, _ := LargestPrimeFactor(v.input)

        if actual != v.output {
            t.Errorf("#%d: LargestPrimeFactor(%d)=%d; expected: %d", i, v.input, actual, v.output)
        }
    }
}

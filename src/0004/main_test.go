package main

import "testing"

type TestCase struct {
    input  int
    output int64
}

func TestLargestPalindromeProduct(t *testing.T) {
    testPairs := []TestCase{
        {1, 9},
        {2, 9009},
        {3, 906609},
    }

    for i, v := range testPairs {
        actual := LargestPalindromeProduct(v.input)

        if actual != v.output {
            t.Errorf("#%d: LargestPalindromeProduct(%d)=%d; expected: %d", i, v.input, actual, v.output)
        }
    }
}

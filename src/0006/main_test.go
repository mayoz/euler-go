package main

import "testing"

type TestCase struct {
    number int
    output int64
}

func TestSumSquareDifference(t *testing.T) {
    testPairs := []TestCase{
        {10, 2640},
        {100, 25164150},
    }

    for i, v := range testPairs {
        actual := SumSquareDifference(v.number)

        if actual != v.output {
            t.Errorf("#%d: SumSquareDifference(%d)=%d; expected: %d", i, v.number, actual, v.output)
        }
    }
}

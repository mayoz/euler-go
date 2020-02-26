package main

func SumSquareDifference(n int) int64 {
    return squareOfSum(n) - sumOfSquare(n)
}

func sumOfSquare(n int) (r int64) {
    for i := 1; i <= n; i++ {
        r += int64(i * i)
    }

    return
}

func squareOfSum(n int) (r int64) {
    for i := 1; i <= n; i++ {
        r += int64(i)
    }

    r *= r

    return
}

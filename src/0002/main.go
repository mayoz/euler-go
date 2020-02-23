package main

func sumEvenFibonacciNumbers(limit int) (sum int) {
    x, y := 1, 2

    for {
        x, y = y, x+y

        if x > limit {
            break
        }

        if (x % 2) == 0 {
            sum += x
        }
    }

    return
}

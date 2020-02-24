package main

func SmallestEvenlyDivisible(n int) int {
    i := n

    for {
        if isDivisible(i, n) {
            return i
        }

        i += n
    }
}

func isDivisible(number, goal int) bool {
    for i := goal; i > 1; i-- {
        if (number % i) != 0 {
            return false
        }
    }

    return true
}

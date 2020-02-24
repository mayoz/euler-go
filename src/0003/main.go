package main

func LargestPrimeFactor(number int64) (prime int64, found bool) {
    for i := int64(2); i <= number; i++ {
        if number%i == 0 {
            if isPrime(i) {
                found = true
                prime = i
            }

            number /= i
            i = 2
        }
    }

    return
}

func isPrime(n int64) bool {
    for i := int64(2); i < n; i++ {
        if n%i == 0 {
            return false
        }
    }

    return true
}

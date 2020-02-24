package main

import "strconv"

func LargestPalindromeProduct(digit int) (largest int64) {
    init, _ := generateNumber(digit)
    product := int64(0)

    for i := init; i > 1; i-- {
        for j := init; j > 1; j-- {
            product = i * j

            if isPalindrome(product) && product > largest {
                largest = product
            }
        }
    }

    return
}

func isPalindrome(n int64) bool {
    str := strconv.FormatInt(n, 10)

    for i, j := 0, len(str)-1; i <= len(str)/2; i++ {
        if str[i] != str[j-i] {
            return false
        }
    }

    return true
}

// Generate a largest number value for given digit. This function is utility,
// it is not related to the solution of the problem.
func generateNumber(digit int) (int64, error) {
    str := ""

    for ; digit > 0; digit-- {
        str += "9"
    }

    return strconv.ParseInt(str, 10, 64)
}

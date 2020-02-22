package main

func sumMultiplies3or5(limit int) (sum int) {
    for i := 1; i < limit; i++ {
        if (i%3 == 0) || (i%5 == 0) {
            sum += i
        }
    }

    return
}

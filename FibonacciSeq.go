package main

import (
"fmt"
"os"
"strconv"
)

//Fibonacci it itirates to the Fibonacci series ith index number in reversal order
func Fibonacci(n int) int {
if n <= 1 {
return n
}
return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
val := os.Args[1]
c, err := strconv.Atoi(val)
if err != nil {
fmt.Println("Expecting interger type argument")
os.Exit(1)
}
for i := 0; i < c; i++ {
fmt.Println(Fibonacci(i)) //Calling Fibonacci here
}
}

// go run filename.go { whatever file name you give} count
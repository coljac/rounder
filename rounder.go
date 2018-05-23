package main

import (
    "fmt"
    "os"
    "strconv"
    "math"
)

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        return
    }
    arg, err := strconv.ParseFloat(args[0], 64)
    if err == nil {
        fmt.Printf("%d\n", int(math.Round(arg)))
    } else {
        fmt.Println(args[0])
    }

}

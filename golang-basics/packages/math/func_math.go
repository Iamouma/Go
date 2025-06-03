package main

import (
    "fmt"
    "math"
)

func main() {
    //Abs returns the absolute value
    a := math.Abs(-5)
    //Round returns the nearest integer, rounding half away from zero.
    b := math.Round(4.544)
    //RoundToEven returns the nearest integer, rounding ties to even.
    d := math.RoundToEven(4.6)
    //Cbrt returns the cube root
    c := math.Cbrt(64)
    //Ceil returns the least integer value greater than or equal to given argument
    e := math.Ceil(4.5)
    //Floor returns the greatest integer value less than or equal to given argument
    f := math.Floor(5.7)
    //Copysign returns a value with the magnitude of first argument and the sign of second argument
    g := math.Copysign(5.6, -5)
    //Max returns the larger of two arguments passed
    h := math.Max(5, 7)
    //Min returns the smaller of two arguments passed
    i := math.Min(5, 7)
    //Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.(x,y are arguments)
    j := math.Mod(10, 3)
    //Nextafter returns the next representable float64 value after x towards y.(x,y are arguments)
    k := math.Nextafter(5.5, 5.8)
    //Trunc returns the integer value of x(x - argument)
    l := math.Trunc(10.998)
    //Remainder returns the IEEE 754 floating-point remainder of x/y.(x,y-arguments)
    m := math.Remainder(11, 9)
    // Pow10 returns 10**n, the base-10 exponential of n.(n-argument)
    n := math.Pow10(4)
    //Pow returns x**y, the base-x exponential of y.(x,y-arguments)
    o := math.Pow(7, 3)
    fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
}

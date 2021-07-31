package main

import (
    "fmt"
    "math"
)

var (
    h = "world"
    i = true
)

func main() {
    variableZeroValue()
    variableInitialValue()
    variableTypeDetection()
    fmt.Println(h, i)
    constDemo()
    enumDemo()
}

func variableZeroValue() {
    var a int
    var s string
    fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
    var a, b int = 3, 4
    var s string = "abc"
    fmt.Println(a, b, s)
}

func variableTypeDetection() {
    var a, b, c, d = 6, 9, true, "hello"
    fmt.Println(a, b, c, d)
}

func constDemo() {
    const fileName = "abc.txt"
    const (
        a, b = 3, 4
    )

    c := int(math.Sqrt(a*a + b*b))
    fmt.Println(fileName, c)
}

func enumDemo() {
    const (
        cpp = iota
        java
        php
        python
    )
    fmt.Println(cpp, java, php, python)

    const (
        b = 1 << (10 * iota)
        kb
        mb
        gb
        tb
        pb
    )
    fmt.Println(b, kb, mb, gb, tb, pb)

}

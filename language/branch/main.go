package main

import (
    "fmt"
    "io/ioutil"
)

func grade(score int) string {
    g := ""
    switch {
    case score < 0 || score > 100:
        panic(fmt.Sprintf("woring score %d \n", score))
    case score < 60:
        return "F"
    case score < 80:
        return "D"
    case score < 90:
        return "B"
    case score <= 100:
        return "A"
    }

    return g

}

func main() {
    const fileName = "./language/branch/abc.txt"
    if contents, err := ioutil.ReadFile(fileName); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s \n", contents)
    }

    fmt.Println(grade(0))
    fmt.Println(grade(70))
    fmt.Println(grade(92))
    // fmt.Println(grade(-20))
    // fmt.Println(grade(120))
}

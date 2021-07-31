package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func convertToBin(n int) string {
    s := ""
    if n == 0 {
        return "0"
    }
    for ; n > 0; n = n / 2 {
        s = strconv.Itoa(n%2) + s
    }

    return s
}

func printFile() {
    const fileName = "./abc.txt"
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}

func main() {
    fmt.Println(convertToBin(5))
    fmt.Println(convertToBin(7))
    fmt.Println(convertToBin(0))
    printFile()
}

package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "hello，年轻人！"
    fmt.Printf("%X\n", s)
    // 中文三个字节，从下面索引能看到
    // 获取之后放到四个字节的rune里
    // ch is rune 等于int32；占用4字节
    for i, ch := range s {
        fmt.Printf("(%d %x)", i, ch)
    }
    // len 获取的是字节长度
    fmt.Println(len(s))

    // 将字节拷贝到新的空间
    fmt.Println(utf8.RuneCountInString(s))

    bytes := []byte(s)
    for len(bytes) > 0 {
        ch, size := utf8.DecodeRune(bytes)
        bytes = bytes[size:]
        fmt.Printf("%c ", ch)
    }

    fmt.Println()
    // 存在空间消耗，重新开辟一块空间来转换。而不是对原有内存数据的重新解码
    for i, ch := range []rune(s) {
        fmt.Printf("[%d-%c] ", i, ch)
    }

}

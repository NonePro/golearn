package main

import (
    "fmt"
    "reflect"
    "runtime"
)

func eval(a, b int, op string) (int, error) {

    switch op {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        return a / b, nil
    default:
        return 0, fmt.Errorf("not supported op %s", op)
    }
}

func div(a, b int) (int, int) {
    return a / b, a % b
}

func apply(op func(a, b int) int, a, b int) int {
    p := reflect.ValueOf(op).Pointer()
    opName := runtime.FuncForPC(p).Name()
    fmt.Printf("Calling function %s  with args (%d, %d) = ", opName, a, b)
    return op(a, b)
}

// 可变参数列表
func sum(nums ...int) int {
    s := 0
    for _, v := range nums {
        s += v
    }
    return s
}

func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    fmt.Println(eval(4, 3, "+"))
    fmt.Println(eval(4, 3, "-"))
    fmt.Println(eval(4, 3, "*"))
    fmt.Println(eval(4, 3, "/"))
    fmt.Println(eval(4, 3, "X"))
    fmt.Println(div(7, 4))
    fmt.Println(apply(func(a, b int) int {
        return a + b
    }, 4, 5))
    fmt.Println(sum(2, 3, 4, 5))
    a, b := 3, 4
    // 定义传递指针，内部依然是复制，go只支持复制,这里是复制的生死指针
    swap(&a, &b)
    fmt.Println(a, b)

}

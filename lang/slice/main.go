package main

import "fmt"

func main() {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Println("arr[2:6r", arr[2:6])
    fmt.Println("arr[:6]", arr[:6])
    s1 := arr[2:]
    fmt.Println("s1", s1)
    s2 := arr[:]
    fmt.Println("s2", s2)

    fmt.Println(arr)
    updateSlice(s1)
    fmt.Println(s1)
    fmt.Println(arr)

    updateSlice(s2)
    fmt.Println(s2)
    fmt.Println(arr)

    fmt.Println("Before ResSlice", s2)
    s2 = s2[2:]
    fmt.Println("After ResSlice", s2)
    s2 = s2[1:3]
    fmt.Println("After ResSlice", s2)

    showExpand()

    showDeclare()
}

func showDeclare() {
    var s []int
    for i := 0; i < 100; i++ {
        s = append(s, 2*i+1)
        fmt.Printf("len(s)=%d cap(s)=%d \n", len(s), cap(s))
    }
    fmt.Println(s)

    s1 := []int{2, 4, 6, 8}
    fmt.Println(s1)

    s2 := make([]int, 16)
    s3 := make([]int, 10, 32)
    fmt.Println(s2, s3)

    fmt.Println("Copying slice")
    copy(s2, s1)
    fmt.Println(s2)

    fmt.Println("Deleting elements from slice")
    // 删除地4个元素 s2 = s2[:3] + s[4:]
    // append 第二个参数开始是单元素,需要后面加三个点
    s2 = append(s2[:3], s2[4:]...)
    fmt.Println(s2)

    fmt.Println("Popping from front")
    front := s2[0]
    s2 = s2[1:]
    fmt.Println(front, s2)

    fmt.Println("Popping from end")
    tail := s2[len(s2)-1]
    s2 = s2[:len(s2)-1]
    fmt.Println(tail, s2)
}

func showExpand() {
    fmt.Println("Extending slice")
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d \n", arr, len(arr), cap(arr))
    s1 := arr[2:6]
    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))
    s2 := s1[3:5]
    fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d \n", s2, len(s2), cap(s2))
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d \n", arr, len(arr), cap(arr))

    s3 := append(s2, 10)
    fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d \n", s3, len(s3), cap(s3))
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d \n", arr, len(arr), cap(arr))

    s4 := append(s3, 11)
    fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d \n", s4, len(s4), cap(s4))
    // 此时已经脱离原来的arr,发生了数组的拷贝
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d \n", arr, len(arr), cap(arr))

    s5 := append(s4, 12)
    fmt.Printf("s5=%v, len(s5)=%d, cap(s5)=%d \n", s5, len(s5), cap(s5))
    fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d \n", arr, len(arr), cap(arr))
}

func updateSlice(slice []int) {
    slice[0] = 100
}

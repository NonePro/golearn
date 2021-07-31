package main

import "fmt"

func main() {
    m1 := map[string]string{
        "name":    "NonePro",
        "course":  "DDD",
        "site":    "demo",
        "quality": "notBad",
    }

    m2 := make(map[string]int)

    var m3 map[int]string

    fmt.Println(m1, m2, m3)

    fmt.Println("Traveling Map")
    // key 在map里是无序的
    for k, v := range m1 {
        fmt.Println(k, v)
    }

    fmt.Println("Getting Values")
    courseName := m1["name"]
    fmt.Println(courseName)

    fmt.Println("访问为定义key", m1["no"], "会返回默默认值,而不是报错")

    fmt.Println("Deleting Values")
    name, ok := m1["name"]
    fmt.Println(name, ok)
    delete(m1, "name")
    name, ok = m1["name"]
    fmt.Println(name, ok)

    // map 的key 使用的hash表，必须可以比较相等
    // 内建类型除了slice,map,function 不可做key,其他都可以
    // struct 不包含上述字段也可以作为key

}

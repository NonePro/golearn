package crawler

import "net/http"

func main() {
    response, err := http.Get("http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun")
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {

    }
}

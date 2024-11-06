package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    apiURL = "https://api.gtja.com/v1/stock/realtime?symbol=000001" // 假设这是国泰君安API的URL
)

func fetchStockPrice() (string, error) {
    resp, err := http.Get(apiURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    data := string(body)
    // 假设返回的数据包含价格字段 "price" ，你需要根据实际的返回格式来解析
    // 示例：假设data包含 "price: 3000"
    return data, nil
}

func main1() {
    for range time.Tick(time.Second) { // 每秒获取一次数据
        price, err := fetchStockPrice()
        if err != nil {
            fmt.Println("获取股票价格失败:", err)
            continue
        }
        fmt.Printf("当前股票价格: %s\n", price)
    }
}

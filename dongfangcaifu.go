package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

const (
    apiURL = "http://push2.eastmoney.com/api/qt/stock/sse?secid=1.000001" // 上证指数的实时数据接口
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
    // 解析数据，提取股票价格等信息
    // 东方财富API返回的数据需要根据实际返回的JSON格式进行解析
    if strings.Contains(data, `"f12"`) {
        // 假设f12是价格字段，具体字段请参考API文档
        parts := strings.Split(data, `"f12":"`)
        if len(parts) > 1 {
            price := strings.Split(parts[1], `"`)
            return price[0], nil // 返回价格
        }
    }

    return "", fmt.Errorf("数据解析失败")
}

func main1() {
    for range time.Tick(time.Second) { // 每秒获取一次数据
        price, err := fetchStockPrice()
        if err != nil {
            fmt.Println("获取股票价格失败:", err)
            continue
        }
        fmt.Printf("当前上证指数: %s\n", price)
    }
}

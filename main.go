package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"github.com/mattn/go-runewidth"
)

// Stock represents the data structure for stock information.
type Stock struct {
	Name          string
	Code          string
	YesterdayClose string
	TodayOpen     string
	Volume        string
	High          string
	Low           string
	ChangeRate    string
	ChangeAmount  string
	CurrentPrice  string
	Buy1          string
	Buy2          string
	Sell1         string
	Sell2         string
}

// fetchStockData retrieves stock data from the URL and returns it as a slice of Stock structs.
func fetchStockData() ([]Stock, error) {
	url := "http://qt.gtimg.cn/q=sh000001,sz002304,sh600519"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode GBK encoded response body to UTF-8
	reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(body), ";")
	stocks := []Stock{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		data := strings.Split(line, "~")
		if len(data) < 41 {
			continue
		}

		// Convert volume to billions with 2 decimal precision
		volume, err := strconv.ParseFloat(data[6], 64)
		if err != nil {
			volume = 0.0
		}
		volumeInBillions := fmt.Sprintf("%.2f", volume/1e8)

		stocks = append(stocks, Stock{
			Name:          data[1],
			Code:          data[2],
			YesterdayClose: data[4],
			TodayOpen:     data[5],
			Volume:        volumeInBillions, // 使用转换后的成交量
			High:          data[33],
			Low:           data[34],
			ChangeRate:    data[32],
			ChangeAmount:  data[31],
			CurrentPrice:  data[3],
			Buy1:          data[9],
			Buy2:          data[11],
			Sell1:         data[19],
			Sell2:         data[21],
		})
	}

	return stocks, nil
}

// padString adjusts the string width to ensure consistent column alignment.
func padString(s string, width int) string {
	return s + strings.Repeat(" ", width-runewidth.StringWidth(s))
}

// printStocks prints stock information in the specified format.
func printStocks(stocks []Stock) {
	headers := []string{"名称", "代码", "昨日收盘", "今日开盘", "成交量(亿)", "今日最高", "今日最低", "涨跌幅", "涨跌额", "当前价格", "买1价", "买2价", "卖1价", "卖2价"}
	columnWidths := []int{10, 8, 10, 10, 12, 10, 10, 8, 8, 10, 8, 8, 8, 8}

	// Print headers
	for i, header := range headers {
		fmt.Print(padString(header, columnWidths[i]))
	}
	fmt.Println()

	// Print stock data
	for _, stock := range stocks {
		values := []string{
			stock.Name, stock.Code, stock.YesterdayClose, stock.TodayOpen, stock.Volume,
			stock.High, stock.Low, stock.ChangeRate, stock.ChangeAmount, stock.CurrentPrice,
			stock.Buy1, stock.Buy2, stock.Sell1, stock.Sell2,
		}
		for i, value := range values {
			fmt.Print(padString(value, columnWidths[i]))
		}
		fmt.Println()
	}
}

//2022年01月02日10:30:20
func GetYMDHMS0() string {
	t := time.Now()
	return fmt.Sprintf("%d年%02d月%02d日%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func main() {
	for {
		stocks, err := fetchStockData()
		if err != nil {
			fmt.Println("Error fetching stock data:", err)
			return
		}

		fmt.Println("更新时间: ", GetYMDHMS0())
		printStocks(stocks)
		time.Sleep(1 * time.Minute)
	}
}

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"sort"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/fatih/structs"
)

const (
	AK  = "YOUR ACCESS KEY"
	SK  = "YOUR SECRET KEY"
	URI = "https://api.biclub.com/api"
)

type OrderPlace struct {
	Number    string `json:"number"`
	OrderType string `json:"orderType"`
	Price     string `json:"price"`
	Symbol    string `json:"symbol"`
	Timestamp int64  `json:"timestamp"`
	AccessKey string `json:"accessKey"`
	Sign      string `json:"sign"`
}

type OrderCancel struct {
	OrderId   string `json:"orderId"`
	Symbol    string `json:"symbol"`
	Timestamp int64  `json:"timestamp"`
	AccessKey string `json:"accessKey"`
	Sign      string `json:"sign"`
}

type OrderBatchCancel struct {
	OrderIdList []string `json:"orderIdList"`
	Symbol      string   `json:"symbol"`
	Timestamp   int64    `json:"timestamp"`
	AccessKey   string   `json:"accessKey"`
	Sign        string   `json:"sign"`
}

type OrderConsign struct {
	OrderId   string `json:"orderId"`
	Symbol    string `json:"symbol"`
	Timestamp int64  `json:"timestamp"`
	AccessKey string `json:"accessKey"`
	Sign      string `json:"sign"`
}

type OrderConsigns struct {
	Symbol    string `json:"symbol"`
	Direct    string `json:"direct"`
	Status    string `json:"status"`
	Size      int32  `json:"size"`
	Timestamp int64  `json:"timestamp"`
	AccessKey string `json:"accessKey"`
	Sign      string `json:"sign"`
}

type BalanceList struct {
	CoinType  string `json:"coinType"`
	Timestamp int64  `json:"timestamp"`
	AccessKey string `json:"accessKey"`
	Sign      string `json:"sign"`
}

type AccountAddress struct {
	CoinTypes string `json:"coinTypes"`
	Timestamp int64  `json:"timestamp"`
	AccessKey string `json:"accessKey"`
	Sign      string `json:"sign"`
}

// Common routines

func makeURL(route string) string {
	u, err := url.Parse(URI)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, route)
	return u.String()
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func mapToSortedString(m map[string]interface{}, sk string) string {
	b := new(bytes.Buffer)

	names := make([]string, 0, len(m)-1)
	for name := range m {
		if name != "Sign" {
			names = append(names, name)
		}
	}
	sort.Strings(names) //sort keys alphabetically
	for _, name := range names {
		v := m[name]
		// if val, ok := v.(string); ok {
		// 	fmt.Fprintf(b, "%s%s", lowerFirst(name), val)
		// 	continue
		// }

		// if val, ok := v.(int64); ok {
		// 	fmt.Fprintf(b, "%s%d", lowerFirst(name), val)
		// 	continue
		// }

		// if val, ok := v.(int32); ok {
		// 	fmt.Fprintf(b, "%s%d", lowerFirst(name), val)
		// 	continue
		// }

		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Slice:
			j, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(b, "%s%s", lowerFirst(name), string(j[:]))

		case reflect.String:
			if val, ok := v.(string); ok {
				fmt.Fprintf(b, "%s%s", lowerFirst(name), val)
			}
		case reflect.Int64:
			if val, ok := v.(int64); ok {
				fmt.Fprintf(b, "%s%d", lowerFirst(name), val)
			}
		case reflect.Int32:
			if val, ok := v.(int32); ok {
				fmt.Fprintf(b, "%s%d", lowerFirst(name), val)
			}
		default:
			panic("Todo: you need to deal with other type")
		}

	}

	fmt.Fprintf(b, "%s", sk)

	return b.String()
}

func sign(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}

func get(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("response Body:", string(respBody))

	return nil
}

func getWithQs(url string, qs map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	q := req.URL.Query()
	for k, v := range qs {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(respBody))

	return nil
}

func post(url string, body []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Accept", "application/json,text/plain, */*")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(respBody))

	return nil
}

// Price APIs

// getMarketSymbols implements the api of symbols of market
// GET /market/common/symbols
func getMarketSymbols(route string) {
	url := makeURL(route)

	if err := get(url); err != nil {
		panic(err)
	}
}

// getMarketTickers implements the api of tickers of market
// GET /market/tickers
func getMarketTickers(route string, qs map[string]string) {
	url := makeURL(route)

	if err := getWithQs(url, qs); err != nil {
		panic(err)
	}
}

// getMarketDepth implements the api of depth of market
// GET /market/depth
func getMarketDepth(route string, qs map[string]string) {
	url := makeURL(route)

	if err := getWithQs(url, qs); err != nil {
		panic(err)
	}
}

// getMarketTrades implements the api of trades of market
// GET /market/trades
func getMarketTrades(route string, qs map[string]string) {
	url := makeURL(route)

	if err := getWithQs(url, qs); err != nil {
		panic(err)
	}
}

// getMarketKline implements the api of kline of market
// GET /market/history/kline
func getMarketKline(route string, qs map[string]string) {
	url := makeURL(route)

	if err := getWithQs(url, qs); err != nil {
		panic(err)
	}
}

// Trade APIs

// postPlace implements the api of order place
// POST /trade/order/orders/place
func postPlace(route string) {
	url := makeURL(route)

	place := &OrderPlace{
		Number:    "1",
		OrderType: "buy-limit",
		Price:     "6530",
		Symbol:    "btc-usdt",
		Timestamp: time.Now().UnixNano() / 1e6,
		AccessKey: AK}

	m := structs.Map(place)
	s := sign(mapToSortedString(m, SK))
	place.Sign = s

	bytesBody, err := json.Marshal(place)
	if err != nil {
		panic(err)
	}

	err = post(url, bytesBody)
	if err != nil {
		panic(err)
	}
}

// postPlaceCancel implements the api of order cancel
// POST /trade/order/orders/cancel
func postPlaceCancel(route string) {
	url := makeURL(route)

	place := &OrderCancel{
		OrderId:   "87e8c3ed-dd11-4522-a968-aba88ae2733b",
		Symbol:    "btc-usdt",
		Timestamp: time.Now().UnixNano() / 1e6,
		AccessKey: AK}

	m := structs.Map(place)
	s := sign(mapToSortedString(m, SK))
	place.Sign = s

	bytesBody, err := json.Marshal(place)
	if err != nil {
		panic(err)
	}

	err = post(url, bytesBody)
	if err != nil {
		panic(err)
	}
}

// postBatchCancel implements the api of batch cancel
// POST /trade/order/orders/batchcancel
func postBatchCancel(route string) {
	url := makeURL(route)

	orderBatchCancel := &OrderBatchCancel{
		OrderIdList: []string{"87068d96-7130-40ed-a4ae-7ad476d4d357", "68206740-7b8f-4232-85d6-e309154aab52", "68206740-7b8f-4232-85d6-e309154aab52"},
		Symbol:      "btc-usdt",
		Timestamp:   time.Now().UnixNano() / 1e6,
		AccessKey:   AK}

	m := structs.Map(orderBatchCancel)
	s := sign(mapToSortedString(m, SK))
	orderBatchCancel.Sign = s

	bytesBody, err := json.Marshal(orderBatchCancel)
	if err != nil {
		panic(err)
	}

	err = post(url, bytesBody)
	if err != nil {
		panic(err)
	}
}

// postConsign implements the api of consign
// POST /trade/order/orders/consign
func postConsign(route string) {
	url := makeURL(route)

	place := &OrderConsign{
		OrderId:   "51de9b3d-11f0-4c46-9aa2-abd389797659",
		Symbol:    "btc-usdt",
		Timestamp: time.Now().UnixNano() / 1e6,
		AccessKey: AK}

	m := structs.Map(place)
	s := sign(mapToSortedString(m, SK))
	place.Sign = s

	bytesBody, err := json.Marshal(place)
	if err != nil {
		panic(err)
	}

	err = post(url, bytesBody)
	if err != nil {
		panic(err)
	}
}

// postOrderConsigns implements the api of consigns
// POST /trade/order/orders/OrderConsigns
func postOrderConsigns(route string) {
	url := makeURL(route)

	place := &OrderConsigns{
		Symbol:    "btc-usdt",
		Direct:    "1",
		Status:    "0,1,2,3",
		Size:      1,
		Timestamp: time.Now().UnixNano() / 1e6,
		AccessKey: AK}

	m := structs.Map(place)
	s := sign(mapToSortedString(m, SK))
	place.Sign = s

	bytesBody, err := json.Marshal(place)
	if err != nil {
		panic(err)
	}

	err = post(url, bytesBody)
	if err != nil {
		panic(err)
	}
}

// postBalanceList implements the api of list of balance
// POST /trade/account/balance/list
func postBalanceList(route string) {
	url := makeURL(route)

	place := &BalanceList{
		CoinType:  "btc",
		Timestamp: time.Now().UnixNano() / 1e6,
		AccessKey: AK}

	m := structs.Map(place)
	s := sign(mapToSortedString(m, SK))
	place.Sign = s

	bytesBody, err := json.Marshal(place)
	if err != nil {
		panic(err)
	}

	err = post(url, bytesBody)
	if err != nil {
		panic(err)
	}
}

// postAccountAddress implements the api of account querying
// POST /trade/account/deposit/address
func postAccountAddress(route string) {
	url := makeURL(route)

	place := &AccountAddress{
		CoinTypes: "bch,btc,usdt",
		Timestamp: time.Now().UnixNano() / 1e6,
		AccessKey: AK}

	m := structs.Map(place)
	s := sign(mapToSortedString(m, SK))
	place.Sign = s

	bytesBody, err := json.Marshal(place)
	if err != nil {
		panic(err)
	}

	err = post(url, bytesBody)
	if err != nil {
		panic(err)
	}
}

func main() {
	// GET requests
	// getMarketSymbols("/market/common/symbols")

	// qs1 := make(map[string]string)
	// qs1["symbol"] = "bz-usdt"
	// getMarketTickers("/market/tickers", qs1)

	// qs2 := make(map[string]string)
	// qs2["symbol"] = "btc-usdt"
	// getMarketDepth("/market/depth", qs2)

	// qs3 := make(map[string]string)
	// qs3["symbol"] = "btc-usdt"
	// qs3["size"] = "10"
	// getMarketTrades("/market/trades", qs3)

	// qs4 := make(map[string]string)
	// qs4["symbol"] = "btc-usdt"
	// qs4["period"] = "30"
	// qs4["size"] = "3"
	// getMarketKline("/market/history/kline", qs4)

	// POST requests
	// postPlace("/trade/order/orders/place")

	// postPlaceCancel("/trade/order/orders/cancel")

	// postConsign("/trade/order/orders/consign")

	// postBatchCancel("/trade/order/orders/batchcancel")

	// postOrderConsigns("/trade/order/orders/consigns")

	// postBalanceList("/trade/account/balance/list")

	// postAccountAddress("/trade/account/deposit/address")
}

package zaif

import (
    "os"
    "time"
    "strconv"
    "crypto/hmac"
    "crypto/sha512"
    "io/ioutil"
    "net/http"
    "encoding/hex"
    "net/url"
    "strings"
	"encoding/json"
	"log"
	"fmt"
)

var key = os.Getenv("ZAIF_KEY")
var secret = os.Getenv("ZAIF_SECRET")

type Depth struct {
    Asks [][]float64 `json:"asks"`
    Bids [][]float64 `json:"bids"`
}

type Trade struct {
	Success int64 `json:success`
	Return struct {
		Received float64 `json:"received"`
		Remains float64 `json:"remains"`
		OrderId int64 `json:"order_id"`
		Funds struct {
			JPY float64 `json:"jpy"`
			BTC float64 `json:"btc"`
			XEM float64 `json:"xem"`
			MONA float64 `json:"mona"`
			BCH float64 `json:"BCH"`
			ETH float64 `json:"ETH"`
			ZAIF float64 `json:"ZAIF"`
		} `json:"funds`
	} `json:return`
}

func GetInfo() string {
	uri := "https://api.zaif.jp/tapi"
	values := url.Values{}
	values.Add("method", "get_info")
	values.Add("nonce", strconv.FormatInt(time.Now().Unix(), 10))

	encodedParams := values.Encode()
	req, _ := http.NewRequest("POST", uri, strings.NewReader(encodedParams))

	hash := hmac.New(sha512.New, []byte(secret))
	hash.Write([]byte(encodedParams))
	signature := hex.EncodeToString(hash.Sum(nil))

	req.Header.Add("Key", key)
	req.Header.Add("Sign", signature)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}

func GetDepth(pair string) Depth {
    uri := "https://api.zaif.jp/api/1/depth/" + pair
    req, _ := http.NewRequest("GET", uri, nil)

    client := new(http.Client)
    resp, _ := client.Do(req)
    defer resp.Body.Close()

    byteArray, _ := ioutil.ReadAll(resp.Body)
	var depth Depth
    if err := json.Unmarshal(byteArray, &depth); err != nil {
        log.Fatal(err)
    }
	// return  strconv.FormatFloat(depth.Asks[0][0], 'f', 6, 64)
	return  depth
    // return string(depth.Ask[0][0])
}

func CreateTrade(pair string, price float64, amount int64) Trade {
	uri := "https://api.zaif.jp/tapi"
	values := url.Values{}
	values.Add("method", "trade")
	values.Add("nonce", strconv.FormatInt(time.Now().Unix(), 10))
	values.Add("action", "bid")
	values.Add("currency_pair", pair)
	values.Add("price", strconv.FormatFloat(price, 'f', 6, 64))
	values.Add("amount", strconv.FormatInt(amount, 10))

	encodedParams := values.Encode()
	req, _ := http.NewRequest("POST", uri, strings.NewReader(encodedParams))

	hash := hmac.New(sha512.New, []byte(secret))
	hash.Write([]byte(encodedParams))
	signature := hex.EncodeToString(hash.Sum(nil))

	req.Header.Add("Key", key)
	req.Header.Add("Sign", signature)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	var trade Trade
    if err := json.Unmarshal(byteArray, &trade); err != nil {
        log.Fatal(err)
    }
	return trade
}

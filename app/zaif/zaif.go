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
)

var key = os.Getenv("ZAIF_KEY")
var secret = os.Getenv("ZAIF_SECRET")

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

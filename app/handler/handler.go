package handler

import (
    "net/http"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/meriy100/ZaifSmoking/app/zaif"
)

func MainPage() echo.HandlerFunc {
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
        return c.String(http.StatusOK, "Hello World")
    }
}

func GetInfo() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, zaif.GetInfo())
	}
}

func GetDepth() echo.HandlerFunc {
	return func(c echo.Context) error {
		pair := c.Param("pair")
		depth := zaif.GetDepth(pair)
		// last := strconv.FormatFloat(depth.Asks[0][0], 'f', 6, 64)
		jsonBytes, _ := json.Marshal(depth)
		return c.String(http.StatusOK, string(jsonBytes))
	}
}

func CreateTrade() echo.HandlerFunc {
	return func(c echo.Context) error {
		pair := c.Param("pair")
		depth := zaif.GetDepth(pair)
		lastBids := depth.Bids[0][0]
		result := zaif.CreateTrade(pair, lastBids, 1)
		jsonBytes, _ := json.Marshal(result)
		return c.String(http.StatusOK, string(jsonBytes))
	}
}


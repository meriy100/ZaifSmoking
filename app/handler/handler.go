package handler

import (
    "net/http"
    "github.com/labstack/echo"
	"github.com/meriy100/zaif/app/zaif"
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

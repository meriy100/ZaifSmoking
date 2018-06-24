package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/meriy100/ZaifSmoking/app/handler"
)

func main() {
    // Echoのインスタンス作る
    e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // ルーティング
    e.GET("/hello", handler.MainPage())
    e.GET("/info", handler.GetInfo())
    e.GET("/depth/:pair", handler.GetDepth())
    e.POST("/trade/:pair", handler.CreateTrade())
    // サーバー起動
    e.Start(":1323")    //ポート番号指定してね
}

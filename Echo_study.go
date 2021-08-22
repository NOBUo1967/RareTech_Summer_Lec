package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

// 何故最初を大文字にしないといけないか
// 最初を大文字にすると外部パッケージからも値を参照できるようになる
// echoパッケージとmainパッケージは別パッケージなので、最初を小文字にしてしまうと値を参照できなくなってしまう
type Turtle struct {
	Name string `json:"name"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
	A float64 `json:"a"`
}

func main() {
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		// err型の値がnilでなければ
		// errを返す
		t := Turtle{"ゼニガメ", 100, 100, 100}
		if err := c.Bind(t); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, t)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

// func main() {
// 	e := echo.New()
// 	t := Turtle("t1", 100, 100, 100)
// 	// 特定のパスを第一引数として設定
// 	// 第二引数には関数を定義
// 	// echo.Context = echoで使われる引数
// 	e.GET("/", func(c echo.Context) error {
// 		// この関数内でデータの取得とデータの加工ができる
// 		// ex) データベースから取得 & 加工できる
// 		message := "TOP PAGE"
// 		return c.String(http.StatusOK, message)
// 	})

	// // ルーティングで別のURLを選択
	// e.GET("/about", func(c echo.Context) error {
	// 	// クエリストリングの書き方
	// 	message := c.QueryParam("ms")
	// 	return c.String(http.StatusOK, message)
	// })
// 	e.Logger.Fatal(e.Start(":1323"))
// }

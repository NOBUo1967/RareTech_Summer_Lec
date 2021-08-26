package main

import "fmt"

// 課題1.
// 2つのfloat64型を引数に対して 「引き算・掛け算・割り算」の計算をし, 
// その値を戻り値として返す関数をそれぞれ作ってください。

func substruct(num1 float64, num2 float64) float64{
	return num1 - num2
}

func multiply(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func divide(num1 float64, num2 float64) float64 {
	return num1 / num2
}

// 課題2.
// 課題1を発展させ, 第2, 3引数にfloat64型の値を入れ, 
// その値に対して第1引数に指定した符号(例: +,-,*,/)に合わせた計算をし, 
// 結果を戻り値として返す関数を作ってください。

// func culc(第一引数 +or-or*or/, 第2引数 num1, 第3引数 num2)

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func calc(symbol string, num1 float64, num2 float64) {
	
	switch symbol {
	case "+":
		add(num1,num2)
	case "-":
		substruct(num1,num2)
	case "*":
		multiply(num1,num2)
	case "/":
		divide(num1,num2)
	default:
		fmt.Println("正しい記号を入力してください")
	}
}

// 課題3.
// 第2回の資料において, Structの2番に書いてあるコードを発展させ以下の要件を満たしたメソッドをそれぞれ作成してください。

// Turtle構造体のaの値を, 引数をもとに足し算するメソッド
// Turtle構造体のnameの値を, 引数をもとに変更するメソッド
// Turtle構造体のnameの値を利用して, msg引数をもとに「Turtleのname : msg」のようなメッセージを出力するメソッド

type Turtle struct {
	name string
	x float64
	y float64
	a float64
	}

	func (t *Turtle) addAngle(a float64) {
		t.a += a
	}

	func (t *Turtle) changeName(name string) {
		t.name = name
	}

	func (t *Turtle) greeting(msg string) {
		fmt.Println(t.name +" : "+ msg)
	}

	// func main() {
	// 	var t1 Turtle = Turtle{"Wartortle", 100,100,100}
	// 	t1.greeting("Good morning.")
	// }

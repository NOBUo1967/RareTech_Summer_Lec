// package main

// パッケージをインポート
// インポートしたパッケージは使用しないとエラーになる
// fmt 文字列の入出力とフォーマットに関する機能を提供するパッケージ
import "fmt"

// 変数宣言
// func main() {
// 	// 宣言した変数は使用しないとエラーになる
// 	var message string = "明記型"
// 	// 方は省略できる
// 	var message = "型省略"
// 	// :=でvarを省略して変数宣言できる
// 	// : = とスペースを開けると認識されない
// 	message := "varも省略"
// }

func main() {
	message := "var省略"
	// Printlen = 改行を含む
	// Print デフォルトフォーマットで出力 改行なし
	// Println 改行を追加する
	// Printf 指定されたフォーマットで出力する
	fmt.Println(message)
	// fmt.Println("Hello,World")
	// fourArithmeticOperations()
	// array()
	// slices()
	// maps()
	swit()
	struc()
}

// constで定数宣言
// 使用していないと警告はでるがエラーにはならない
// 同じスコープで同名の変数・定数を定義すると重複でエラーになる
// 上記のmain関数内でconst messageとするとエラーになる
// const message = "test"

func fourArithmeticOperations() {
	var num1 int = 1
	var num2 int = 2

	fmt.Println(num1 + num2)

	fmt.Println(num1 - num2)

	fmt.Println(num1 * num2)

	fmt.Println(num2 / num1)

	fmt.Println(num2 % num1)
}

func array() {
	// goの配列は固定長ため最初に宣言した配列の要素の数を超えることはできない
	// var 変数名[長さ]型 = [大きさ]型{初期値1,初期値n}
	// 変数名 := []型{初期値...} 要素数を省略しているが{}に書いた要素数を自動でカウントしているため、明記した場合と違いはない
	// var nums[3] = [3]int{1,4} 要素数が足りていないが足りないところには0が入る
	var nums [3]int = [3]int{}
	nums[0] = 1
	nums[1] = 5
	// nums[3]は配列の要素数が宣言より多くなるためエラーになる
	// nums[3] = 11
	fmt.Println(nums)
	// 要素の数が足りていないため0が入っている
	// => [1, 5, 0]
	// len(配列名)で要素の数を返す
	fmt.Println(len(nums))
}

func slices() {
	// 配列と違い可変長のため追加が可能
	// 配列を参照できるため配列から値を切り出すことが可能
	// var 変数名 []型 = []型{初期値1,初期値n}
	var nums []int = []int{}
	nums = append(nums, 1)
	nums = append(nums, 5)
	nums = append(nums, 9)
	fmt.Println(nums)
}

func maps() {
	// 連想配列
	//  変数名 := map[keyの型]valueの型{}
	// 初期値を指定する宣言方法
	status := map[string]int{
		"hp": 1000,
		// 改行する際は,が必須
		// } 閉じカッコを改行するときも,が必要
		"mp": 50,
		}

		fmt.Println(status)
		// 要素の挿入
		status["atk"] = 50
		status["def"] = 100
		// Printlnで出力するとアルファベット順に出力される
		// key = defで挿入するとatk,def,hpの順で出力される
		// key = zet で挿入するとatk,hp,zetの順で出力される
		fmt.Println(status)
		// 要素の削除
		delete(status, "mp")
		fmt.Println(status)
}

func swit() {
	num := 1

	switch num {
	case 1:
		fmt.Println("1のとき")
		// breakは必要ない
		// 次のcaseの処理もしたい場合は
		fallthrough
		// を記載する
	case 2:
		fmt.Println("2のとき")
	default:
		fmt.Println("それ以外")
	}
	// swit() => 1のとき、2のとき となる
}

// goの繰り返し文はfor文のみ
// while や do while はない
// do {処理} while(条件) 
// => do {}の処理を1回は実行し、その後に条件を判定。trueなら処理を繰り返す
// コンポジット型にはrangeが使える for eachみたいな
func rep() {
	num := 1
	
	// 普通の書き方
	for i := 0 ; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0 ; i < 5; i++ {
		if i == 1 {
			continue
			// continueはその回のみskipされる
			// for文にはbreakがある。breakはfor文そのものを終了させる。
		}
		fmt.Println(i)
	}

	// while文みたいな
	for num < 5 {
		fmt.Println(num)
		num++
	}

	// 配列をfor文で回す
	nums := [2]int{1,2}
	// for index, 要素 := range 配列など {}
	// indexや要素は宣言したら使わないといけない
	// 使いたくない場合は_にする
	// for _, num := range nums {}
	for i, num := range nums {
		fmt.Println(i)
		fmt.Println(num)
	}
}

// structの定義
type Turtle struct {
	name string
	x float64
	y float64
	a float64
}

// structの関数の定義
// 関数の定義はstructの外で定義する
// (t *Turtle) 
// *Turtle = struct Turtleで使うという意味
// tはTurtleの別名 t.x = Turtle.x
func (t *Turtle) move(x float64, y float64) {
	t.x += x
	t.y += y
}

func struc() {
	// structの宣言の仕方
	var t1 Turtle = Turtle{"師匠", 1000, 5, 180.5}
	var t2 Turtle = Turtle{"弟子", 10, 250, 270.3}
	fmt.Println(t1)
	fmt.Println(t2)

	t1.move(10, -10)
	fmt.Println(t1)
}



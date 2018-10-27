package main //每个 Go 应用程序都包含一个名为 main 的包
import (
	"fmt"
)
const (
	Unknown = 0
	Female = 1
	Male = 2
)
const ( //iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	a = iota
	b = iota
	c = iota
)
// 和java js 差不多

func main() { // 必须 为main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")

	var time int = 5
	time = time + time
	fmt.Println(time)

	//var a int = 10
	//fmt.Println(a)
	//var b = "fdsafdsa"  //可以自动判断 变量 类型
	//fmt.Println(b)
	//c := 10 // 可以 不申明 但是要有 ：=
	//fmt.Println(c)
	//类型相同多个变量, 非全局变量

	//var x, y int
	//var (  // 这种因式分解关键字的写法一般用于声明全局变量
	//	a int
	//	b bool
	//)
	//
	//var c, d int = 1, 2
	//var e, f = 123, "hello"
	//
	//g, h := 123, "hello"
	//println(x, y, a, b, c, d, e, f, g, h)


	print(len("fdsafds"))

	const  fsda  = "fsda";
	var a int = 21
	var b int = 10

	if( a == b ) {
		fmt.Printf("第一行 - a 等于 b\n" )
	} else {
		fmt.Printf("第一行 - a 不等于 b\n" )
	}
	if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n" )
	} else {
		fmt.Printf("第二行 - a 不小于 b\n" )
	}

	if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n" )
	} else {
		fmt.Printf("第三行 - a 不大于 b\n" )
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于 b\n" )
	}
	if ( b >= a ) {
		fmt.Printf("第五行 - b 大于等于 a\n" )
	}


	var var1 = 55
	switch var1 {
	case 44:
		fmt.Printf("44")
	case 55:
		fmt.Printf("55")
	default:
		fmt.Printf("default")
	}

	/*switch stype*/


	var x interface{}
	switch x.(type){
	case int:
		fmt.Printf("int")
	case string:
		fmt.Printf("string")
		/* 你可以定义任意个数的case */
	default: /* 可选 */
		fmt.Printf("any")
	}


	//select










}

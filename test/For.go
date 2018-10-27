package main

import "fmt"

func main() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5} //数组

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if (i%j == 0) {
				break; // 如果发现因子，则不是素数
			}
		}
		if (j > (i / j)) {
			fmt.Printf("%d  是素数\n", i);
		}
	}

	/* 定义局部变量 */
	var a1 int = 10

	/* 循环 */
	LOOP:for a1 < 20 {
		if a1 == 15 {
			/* 跳过迭代 */
			a1 = a1 + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a1)
		a1++
	}

	var dsa =  testFun()
	fmt.Printf("dsa的值为 : %d\n", dsa)

	aq, bq := swap("Mahesh", "Kumar")
	fmt.Println(aq, bq)

}

func testFun()int {

	return 65
}

func swap(x, y string) (string, string) {
	return y, x
}

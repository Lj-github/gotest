package main

import "fmt"

func main() {

	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}


	balance[4] = 50.0

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}


	//Go 语言切片(Slice)   类似  python

	var numbersd = make([]int,3,5)

	printSlice(numbersd)


	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)


	//copy

	var numbers2 []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers2 = append(numbers2, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers2 = append(numbers2, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers2 = append(numbers2, 2,3,4)
	printSlice(numbers2)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers3 := make([]int, len(numbers2), (cap(numbers2))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers3,numbers2)
	printSlice(numbers3)



}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
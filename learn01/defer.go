package main

import "fmt"

func main() {

	//入栈 出栈 先打印end2 后end1 与入栈顺序相反
	defer fmt.Println("defer main end1")
	defer fmt.Println("defer main end2")

	defer funcA()
	defer funcB()
	defer funcC()

	fmt.Println("main hello go1")
	fmt.Println("main hello go2")

	// return在defer前调用
	deferAndReturnFunc()
}

func funcA() {
	fmt.Println("A")
}
func funcB() {
	fmt.Println("B")
}
func funcC() {
	fmt.Println("C")
}

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func deferAndReturnFunc() int {
	defer deferFunc()
	return returnFunc()
}

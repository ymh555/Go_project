package main

import (
	"fmt"
	"module1/package1"
	"module1/package2"
)

func main() {
	fmt.Println("hello world")
	// 直接调用main包内的函数
	M2_say()
	M3_say()

	// 调用package1内的函数
	package1.Xxx1_say()
	package1.Xxx2_say()
	package1.Xxx3_say()

	// 调用package2内的函数
	package2.Yyy1_say()
	package2.Yyy2_say()
	package2.Yyy3_say()

}

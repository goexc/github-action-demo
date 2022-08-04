package main

import "fmt"

func main() {
	say := Cat()
	fmt.Println(say)
}

func Cat() string {
	//格式化输出
	return fmt.Sprintf("%s~~~", "Miao")
}

package main

import "fmt"

func main() {
	say := Cat()
	fmt.Println(say)
}

func Cat() string {
	var password = "12345"
	fmt.Println(password)
	//格式化输出
	return fmt.Sprintf("%s~~~", "Miao")
}

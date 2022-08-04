package main

import "fmt"

func main() {
	say := Cat()
	fmt.Println(say)
}

func Cat() string {
	return fmt.Sprintf("%s~~~", "Miao")
}

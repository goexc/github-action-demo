package main

import "testing"

func TestCat(t *testing.T) {
	say := Cat()
	if say != "Miao~~" {
		t.Errorf("Cat say %s", say)
	}
}

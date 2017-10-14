package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
	"reflect"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		ch <- t.Value
		if t.Left != nil {
			Walk(t.Left, ch)
		}
		if t.Right != nil {
			Walk(t.Right, ch)
		}
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool { 
	c1, c2:= make(chan int, 10), make(chan int, 10)
	go Walk(t1, c1)
	go Walk(t2, c2)
	
	s1, s2 := make([]int, 10, 10), make([]int, 10, 10)
	func() {
		for i := 0; i < 10; i++ {
			s1[i] = <- c1
		}
		sort.Ints(s1)
	}()
	func() {
		for i := 0; i < 10; i++ {
			s2[i] = <- c2
		}
		sort.Ints(s2)
	}()
	
	return reflect.DeepEqual(s1, s2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))

}

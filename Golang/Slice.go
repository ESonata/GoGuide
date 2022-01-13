package main

import "fmt"

func Xianjiang(a *[]int) {
	*a = append(*a, 1)
}
func Xianjiang2(a []int) {
	a = append(a, 2)
}

type Door bool

func (d *Door) Open() {
	fmt.Println("Open the door")
}

func (d *Door) Close() {
	fmt.Println("Close the door")
}
func t(m chan struct{}) {

	m <- struct{}{}
}
func main() {
	a := make([]int, 0)
	b := make([]int, 0)
	Xianjiang(&a)
	fmt.Println(a)
	Xianjiang2(b)
	fmt.Println(b)
	m := make(chan struct{})
	select {
	case m <- struct{}{}:
		fmt.Println("!1")
	default:
		fmt.Println("111")
	}
}

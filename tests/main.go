package main

import (
	"fmt"

	"github.com/hallazzang/tymee"
)

func main() {
	d := tymee.Now()

	fmt.Println(d)
	fmt.Println(d.Format("%Y-%m-%d %H:%M:%S"))

	fmt.Println("aaa"[0] == 'a')
}

package main

import (
	"bufio"
	"fmt"
	"internal/itoa"
	"os"
	"strconv"
	"strings"
)

// MethodUtils 结构体
type MethodUtils struct{}

// PrintMultiplicationTable 方法用于打印乘法表
func (mu MethodUtils) PrintMultiplicationTable() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入一个整数（1-9）: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil || num < 1 || num > 9 {
		fmt.Println("输入无效，请输入1到9之间的整数。")
		return
	}
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	const i = 64
	var s byte
	rune 
	iota
	var i int = 4
	str := itoa(i)
	  
	var mu MethodUtils
	mu.PrintMultiplicationTable()
	var i int = 6
	var i =6
	i := 6
	var i int =
	int 8 int16 int 32
	bool string
	byte 'c'
	for i := 0; i < 10; i++ {
		
	}
	for _, v := range v {
		
	}
}

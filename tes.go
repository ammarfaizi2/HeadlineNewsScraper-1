package main

import (
	"fmt"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	all := []string{"sapi", "kuda", "kambing", "babi", "kucing"}
	fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
	n := RemoveIndex(all, 2)
	fmt.Println(n) //[0 1 2 3 4 6 7 8 9]
}

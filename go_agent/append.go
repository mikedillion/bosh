package main

import (
	"fmt"
	"strconv"
)	

func main() {
	var stdout string
	for i := 0; i < 100; i++ {
		stdout = fmt.Sprintf("%s%s\n", stdout, strconv.Itoa(i))
	}
	fmt.Println("%s", stdout)
}

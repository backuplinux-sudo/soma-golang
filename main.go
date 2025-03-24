package main

import (
	"fmt"
	"strconv"
	"os"
	"github.com/softchris/math"

)

func soma(num1, num2 int) int {
	return math.Add(num1, num2)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Por favor, forneça dois números inteiros como argumento.")
		return
	}

	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])
	fmt.Println(soma(num1,num2))
}

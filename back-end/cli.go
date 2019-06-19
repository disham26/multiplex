package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)
	txtInt, _ := strconv.Atoi(text)
	fmt.Println(txtInt)
	fmt.Print("Enter text: ")
	text1, _ := reader.ReadString('\n')
	fmt.Println(text1)

}

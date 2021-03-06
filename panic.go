package main

import "runtime"
import "os"
import "fmt"

func main() {
	printStack()
	fmt.Println("++++++++++++++++++")
	defer printStack()
	f(3)

}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(i int) {
	defer fmt.Println(i, 5/i)
	f(i - 1)
}

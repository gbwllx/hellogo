package main

//go引入的是包，不是具体的文件
import (
	"os"
	"strconv"
	"fmt"

	"ch2new"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := ch2new.Fahrenheit(t)
		c := ch2new.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, ch2new.FToC(f), c, ch2new.CToF(c))
	}
}

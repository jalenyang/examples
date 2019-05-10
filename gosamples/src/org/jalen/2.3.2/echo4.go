package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println("args",os.Args)
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	a:=new(int64)
	fmt.Println(a)
}

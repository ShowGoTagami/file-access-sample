package main

import (
	"os"
	"fmt"
)

func main() {
	wt := func(f *os.File, s string) {
		_, er := f.WriteString(s + "\n")
		if er != nil {
			fmt.Println(er)
			f.Close()
			return
		}
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Println("*** start ***")
	wt(f, "*** start ***")
	wt(f, "write text!")
	wt(f, "*** end ***")
	fmt.Println("*** end ***")
	er = f.Close()
	if er != nil {
		fmt.Println(er)
	}
}

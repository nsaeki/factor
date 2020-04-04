package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-isatty"
)

func main() {
	var multiLine bool
	flag.BoolVar(&multiLine, "1", false, "print each factor per line")
	flag.Parse()
	multiLine = multiLine || !isatty.IsTerminal(os.Stdout.Fd())

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	var n int64
	n, err := strconv.ParseInt(flag.Args()[0], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	ret := []int64{}
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			n /= i
			ret = append(ret, i)
		}
	}

	if multiLine {
		for _, v := range ret {
			fmt.Println(v)
		}
	} else {
		fmt.Println(strings.Trim(fmt.Sprint(ret), "[]"))
	}

}

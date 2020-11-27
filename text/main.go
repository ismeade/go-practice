package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	d, _ := time.ParseDuration("3m38")
	for i := 0; i < 10; i++ {
		fmt.Fprintf(tw, format, "tt", "art", "alb", i, d)
	}
	tw.Flush()
}

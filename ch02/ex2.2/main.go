package main

import (
	"fmt"
	"github.com/deputatovn/gopl-solutions/ch02/ex2.2/itomconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := itomconv.Celsius(val)
		f := itomconv.Fahrenheit(val)
		ft := itomconv.Foot(val)
		m := itomconv.Meter(val)
		lb := itomconv.Pound(val)
		kg := itomconv.Kilogram(val)

		fmt.Printf("%s = %s, %s = %s\n", f, itomconv.FToC(f), c, itomconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", ft, itomconv.FToM(ft), m, itomconv.MToF(m))
		fmt.Printf("%s = %s, %s = %s\n", lb, itomconv.PToK(lb), kg, itomconv.KToP(kg))
	}
}

package main

import (
	"strings"
)

var args = []string {"first", "second", "third", "fourth"}

func Echo1(){
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

func Echo2(){
	s, sep := "", " "
	for _, arg := range args{
		s += sep + arg
		sep = " "
	}
}

func Echo3(){
	strings.Join(args, " ")
}

func main() {
	Echo1()
	Echo2()
	Echo3()
}
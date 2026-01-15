package main

import "os"

type Usecase interface {
	Execute()
}

var registry = map[string]Usecase{
	"1": &BadImplementation{},
	"2": &GoodImplementation{},
}

func main() {
	args := os.Args[1:]
	registry[args[0]].Execute()
}

package main

type Usecase interface {
	Execute()
}

var registry = map[string]Usecase{
	"1": &BadImplementation{},
	"2": &GoodImplementation{},
}

package main

import (
	"fmt"
)

type FooBar struct {
	n              int
	streamFooToBar chan struct{}
	streamBarToFoo chan struct{}
	streamEnd      chan struct{}
}

func (this *FooBar) Foo(printFoo func()) {
	for i := 0; i < this.n; {
		<-this.streamBarToFoo
		printFoo()
		i++
		this.streamFooToBar <- struct{}{}
	}

	<-this.streamBarToFoo
}

func (this *FooBar) Bar(printBar func()) {
	for i := 0; i < this.n; {
		<-this.streamFooToBar
		printBar()
		i++
		this.streamBarToFoo <- struct{}{}
	}

	this.streamEnd <- struct{}{}
}

func main() {
	var PrintFooBar = func(times int) {
		fooBar := &FooBar{
			n:              times,
			streamFooToBar: make(chan struct{}),
			streamBarToFoo: make(chan struct{}),
			streamEnd:      make(chan struct{}),
		}

		go fooBar.Foo(func() { fmt.Printf("Foo") })
		go fooBar.Bar(func() { fmt.Printf("Bar ") })
		fooBar.streamBarToFoo <- struct{}{}
		<-fooBar.streamEnd
		fmt.Println()
	}

	testCase := []int{0, 1, 2, 3, 4, 5, 7}

	for _, repeat := range testCase {
		fmt.Printf("Repeat %d: ", repeat)
		PrintFooBar(repeat)
	}
}


/*
Output

Repeat 0: 
Repeat 1: FooBar 
Repeat 2: FooBar FooBar 
Repeat 3: FooBar FooBar FooBar 
Repeat 4: FooBar FooBar FooBar FooBar 
Repeat 5: FooBar FooBar FooBar FooBar FooBar 
Repeat 7: FooBar FooBar FooBar FooBar FooBar FooBar FooBar 

*/
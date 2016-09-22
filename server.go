package main

import "fmt"

type First struct {
	Name  string
	Chan  chan string
	Close chan bool
}

func NewFirst(name string) *First {
	return &First{
		Name:  name,
		Chan:  make(chan string),
		Close: make(chan bool),
	}

}

type Second struct {
	Chan chan bool
}

//
func NewSecond() *Second {
	return &Second{
		Chan: make(chan bool),
	}
}

// Test comment
func Test(f *First, s *Second) {
	for msg := range f.Chan {
		_ = "breakpoint"
		fmt.Println(msg)
		f.Name = msg
		s.Chan <- true
	}
}

// Test2 comment
func Test2(f *First, s *Second, stop chan struct{}) {
	msgs := []string{"Name1", "Name2", "Name3", "Name4"}
	for _, msg := range msgs {
		f.Chan <- msg
		<-s.Chan
		if f.Name == msg {
			fmt.Println("Name CHANGED!")
		}
	}
	close(stop)

}

func main() {
	first := NewFirst("Name from main")
	second := NewSecond()
	stop := make(chan struct{})
	go Test(first, second)
	go Test2(first, second, stop)
	<-stop
	fmt.Println("END")
}

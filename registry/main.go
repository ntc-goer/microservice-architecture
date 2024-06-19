package main

import (
	"fmt"
	"time"
)

// const WORKER_1 = "WORKER_1"
// const WORKER_2 = "WORKER_2"
//
// // WORKER_1 -> Event => WORKER_2_QUEUE =>
//
//	type Worker interface {
//		Handle() error
//	}
//
// type Worker1 struct{}
//
//	func (wk *Worker1) Handle() {
//		constraints.Ordered()
//		fmt.Println("Worker 1")
//	}
//
//	func NewWorker1() *Worker1 {
//		return &Worker1{}
//	}
//
// type Worker2 struct{}
//
//	func NewWorker2() *Worker2 {
//		return &Worker2{}
//	}
func main() {
	channel := make(chan string, 5)

	for i := 0; i < 3; i++ {
		go func(testChan <-chan string, index int) {
			for {
				result, isAlive := <-testChan
				if !isAlive {
					break
				}
				fmt.Println(fmt.Sprintf("Receive from worker %d value %s --- Current Bugffer %d", index, result, len(channel)))
				time.Sleep(50 * time.Millisecond)
			}
		}(channel, i+1)
	}
	time.Sleep(3 * time.Second)
	for i := 0; i < 200; i++ {
		channel <- fmt.Sprintf("Number is %d", i)
	}

	time.Sleep(5 * time.Second)
}

// interface in struct
//type BarI interface {
//	Print()
//}
//
//type Bar struct {
//	Value string
//}
//
//func (b *Bar) Print() {
//	fmt.Println("Print Bar Value")
//}
//
//type Foo struct {
//	BarI
//}
//
//func (f *Foo) Print() {
//	fmt.Println("Print Foo Value")
//}
//
//func view(bi BarI) {
//	bi.Print()
//}
//
//func TestInterfaceInStruct(t *testing.T) {
//	// (1) When initializing the struct with properties as an interface , need to pass the struct implemented that interface when initializing the struct
//	foo := Foo{
//		&Bar{},
//	}
//
//	// (2): 2 way to access nested embedded method of a struct
//	//1. Directly access to method of nested struct
//	foo.Print()
//	//2. Access method via struct name
//	foo.BarI.Print()
//
//	// (3) Struct contain interface implement interface too
//	view(&foo)
//
//	assert.Equal(t, nil, nil)
//}

// implement interface
//type FooI interface {
//	Print()
//}

// This is a Foo struct implementing FooI , buw when initializing this struct , we have to pass the definition for FooI properties
//type Foo struct {
//	FooI
//}
//
//type Bar struct{}
//
//func (b *Bar) Print() {
//	fmt.Println("Print Bar Value")
//}
//
//type Tar struct{}
//
//func (t *Tar) Print() {
//	fmt.Println("Print Tar Value")
//}

//type Foo struct {
//}
//
//func (f *Foo) Print() {
//	fmt.Println("Print Foo Value")
//}
//
//func view(f FooI) {
//	f.Print()
//}
//func TestInterface(t *testing.T) {
//	// Use case: When want to dynamic create a struct dynamic implement interface
//	//foo1 := Foo{
//	//	FooI: &Bar{},
//	//}
//	//view(&foo1)
//	//
//	//foo2 := Foo{
//	//	FooI: &Tar{},
//	//}
//	//view(&foo2)
//
//	// Use case: When want to create a struct with already implemented interface's  method before
//	foo := Foo{}
//	view(&foo)
//}

package main

import "fmt"

var result []string

func CreateBinary(preStr string) {
	binaryNumList := []string{"0", "1"}
	if len(preStr) == 3 {
		result = append(result, preStr)
		return
	}
	for _, s := range binaryNumList {
		CreateBinary(preStr + s)
	}
}
func main() {
	// 000 , 001 , 010 , 011
	CreateBinary("")
	fmt.Println(result)
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

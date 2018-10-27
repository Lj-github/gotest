package main

import "fmt"

// go 语言 中的类  就是通过 type   struct  基础  类似 C？

type Poem struct {
	Title  string
	Author string
	intro  string
}

type FlyingCreature struct {
	Poem  // 继承
	WingSpan int
}
func (poem *Poem) publish() {
	fmt.Println("poem publish")
}

func (poem Poem) trr() {
	fmt.Println("poem trr")
	fmt.Println("poem trr" , poem.Title)
}

func (poem FlyingCreature) dump() {
	fmt.Println("poem trr")
	fmt.Println("poem trr" , poem.Title)
	fmt.Println("poem trr" , poem.WingSpan)
}


type Fooer interface {
	Foo1()
	Foo2()
	Foo3()
}

type Foo struct {
}

func (f Foo) Foo1() {
	fmt.Println("Foo1() here")
}

func (f Foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f Foo) Foo3() {
	fmt.Println("Foo3() here")
}


func main() {


	var d =  Poem{Title:"fdsa"}
	fmt.Println(d.Title)
	fmt.Println(d.intro)
	//相当于面向对象
	Poem.trr(d)



}
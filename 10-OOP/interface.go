package main

import "fmt"

// 实现了 interface 的所有方法的类，就实现了这个 interface。并且可以使用绑定在这个 interface 上的方法
// 用 interface 才能实现多态
// interface 定义抽象方法
// 本质是一个父类指针，指向了它本身，同时还指向了要实现的方法。子类要继承
// 多态的基本要素： 有父类，有子类。父类类型的变量指向之类的具体数据变量。
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 把 interface 的方法实现即为实现了该 interface
type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("The cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("The dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return "dog"
}

// 传递 AnimalIF 的不同实现对象，得到的结果不同，这就是多态的现象
func ShowAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}

func main() {
/*	cat := Cat{Color: "yellow"}
	cat.Sleep()
	catColor := cat.GetColor()
	fmt.Println(catColor)
	fmt.Println(cat.GetType())

	fmt.Println("===========")
	ShowAnimal(&cat)

	fmt.Println("===========")
	dog := Dog{Color: "white"}
	ShowAnimal(&dog)*/

/*	var animal AnimalIF // interface 的数据类型，父类指针
	animal = &Cat{Color: "Black"} // interface 本质是一个指针，所以要赋值 Cat 的指针
	animal.Sleep() // 调用 Cat 的 Sleep

	animal = &Dog{Color: "yellow"}
	animal.Sleep() //  调用 Dog 的 Sleep，多态的现象*/

	cat := Cat{Color: "white"}

	dog := Dog{Color: "black"}

	ShowAnimal(&cat)
	ShowAnimal(&dog)
}

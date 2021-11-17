package main

import "fmt"

type Human struct {
	Name string
	Sex string
}

func (this *Human) Eat() {
	fmt.Println("Human is eating ...")
}

func (this *Human) Work() {
	fmt.Println("Human is working ...")
}

// ======
type SuperMan struct {
	Human // 该类继承了 Human 类的方法
	Level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan is eating ...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan is flying ...")
}

func (this *SuperMan) Print() {
	fmt.Println(this.Name, this.Sex, this.Level)
}

func main() {
	human := Human{Name: "Jack", Sex: "female"}
	human.Eat()
	human.Work()

	//superMan := SuperMan{Human: human, Level: 1}
	//superMan := SuperMan{Human: Human{Name: "Michael", Sex: "male"}, Level: 1}
	var superMan SuperMan
	superMan.Name = "Michael"
	superMan.Sex = "male"
	superMan.Level = 1
	// 子类方法
	superMan.Eat()
	// 父类方法
	superMan.Work()
	// 子类方法
	superMan.Fly()
	superMan.Print()
}

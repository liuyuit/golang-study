package main

import "fmt"

// 类名首字母大写表示该类是可以对外使用的
// 类属性首字母大写表示该属性可以对外访问，否则只能在当前包的内部访问。
// go 语言的封装是针对包来封装的
type Hero struct {
	Name string
	Ad int
	Level int
}

/*func (this Hero) Show() {
	fmt.Println(this.Name, this.Ad, this.Level)
}

// (this Hero) 表示该方法绑定到 Hero 这个结构体， this 表示当前对象
func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	// this 是调用该方法的对象的一个拷贝，对它的修改不影响原对象
	this.Name = newName
}*/

// 首字母大写表示其它包可以访问
func (this *Hero) Show() {
	fmt.Println(this.Name, this.Ad, this.Level)
}

// (this Hero) 表示该方法绑定到 Hero 这个结构体， this 表示当前对象
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	// this 是调用该方法的对象的指针，对它的修改会影响原对象
	this.Name = newName
}

func main() {
	hero := Hero{Name : "Jack", Ad : 100, Level: 1}
	hero.Show()
	hero.SetName("Micheal")
	hero.Show()
}
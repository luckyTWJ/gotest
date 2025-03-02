package main

import "fmt"

type Hero struct {
	//如果类的属性首字母大写 表示该属性对外能够访问，否则只能类内部访问 属性也一样
	name string
	age  int
}

/*func (this Hero) GetName() string {
	return this.name
}

func (this Hero) SetName(newName string) {
	//this 调用该方法对象的一个副本
	this.name = newName
}

func (this Hero) Show() {
	fmt.Println("name=", this.name)
	fmt.Println("age=", this.age)
}*/

func (this *Hero) GetName() string {
	return this.name
}

func (this *Hero) SetName(newName string) {
	//this 调用该方法对象的一个副本
	this.name = newName
}

func (this *Hero) Show() {
	fmt.Println("-----------------")
	fmt.Println("name=", this.name)
	fmt.Println("age=", this.age)
}
func main() {
	hero := Hero{name: "hero", age: 18}
	hero.GetName()
	hero.Show()
	hero.SetName("hero2")
	hero.Show()

}

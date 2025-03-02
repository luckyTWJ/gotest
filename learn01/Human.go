package main

import "fmt"

type Human struct {
	name string
	age  int
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat...")
}
func (this *Human) Run() {
	fmt.Println("Human run...")
}
func (this *Human) GetName() string {
	return this.name
}

type SubHuman struct {
	Human //匿名字段 继承Human
	level int
}

func (this *SubHuman) Play() {
	fmt.Println("SubHuman play...")
}
func (this *SubHuman) Eat() {
	fmt.Println("SubHuman eat...")
}
func main() {

	h := Human{"zhang3", 18, "man"}
	h.Eat()
	h.Run()
	fmt.Println("------------sub human-----------")
	//定义子类
	sh := SubHuman{Human{"li4", 19, "woman"}, 100}
	sh.Eat()
	sh.Run()
	sh.Play()
	fmt.Println("------------anonymous-----------")
	var shh SubHuman
	shh.name = "wang5"
	shh.age = 20
	shh.sex = "man"
	shh.Eat()
	shh.Run()
	shh.Play()

}

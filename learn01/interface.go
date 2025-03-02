package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Run()
	Eat()
	GetColor() string
	GetType() string
	Sleep()
}

// 一个接口实现2个类 cat dog
type DogIF interface {
	AnimalIF
	Jump()
}
type Cat struct {
	color   string
	typeStr string
}
type Dog struct {
	color   string
	typeStr string
}

func (this *Cat) Run() {
	fmt.Println("Cat is run....")
}
func (this *Cat) Eat() {
	fmt.Println("Cat is eat....")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return this.typeStr
}
func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep....")
}
func (this *Dog) Jump() {
	fmt.Println("Dog is jump....")
}
func (this *Dog) Eat() {
	fmt.Println("Dog is eat....")
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep....")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return this.typeStr
}
func (this *Dog) GetName() string {
	return "Dog"
}
func (this *Dog) Run() {
	fmt.Println("Dog is run....")

}
func main() {
	//cat := Cat{color: "white", typeStr: "cat"}
	//dog := Dog{color: "black", typeStr: "dog"}
	var animal AnimalIF
	animal = &Dog{color: "black", typeStr: "dog"}
	animal.Run()
	animal.Eat()

	animal = &Cat{color: "white", typeStr: "cat"}
	animal.Run()
	animal.Eat()

}

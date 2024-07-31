package main

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	// TODO: 实现接口
	var animal Animal
	dog := Dog{"Buddy"}
	//cat := Cat{"Whiskers"}
	animal = dog // 猫是Animal接口的实现
	//animal = cat // 狗也是Animal接口的实现
	println(animal.Speak())
}

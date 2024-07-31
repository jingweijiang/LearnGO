package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Height float32
	Sex    bool
}

func main() {

	a := Human{
		Name:   "alex",
		Age:    20,
		Height: 1.73,
		Sex:    false,
	}

	fmt.Println(a)
	fmt.Printf("%d %s %.2f %t \n", a.Age, a.Name, a.Height, a.Sex)
	fmt.Printf("%v \n", a)
	fmt.Printf("%+v \n", a)
	fmt.Printf("%#v \n", a)

}

package main

import "fmt"

func main() {

	ch := make(chan int)
	<-ch
	fmt.Println("Deadlock!")
	//go func() {
	//	<-ch
	//}()
	//
	//go func() {
	//	time.Sleep(1 * time.Second)
	//}()
	//
	//<-ch

}

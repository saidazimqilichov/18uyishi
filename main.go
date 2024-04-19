package main

import "fmt"
func factorial(n int,ch chan int){
	kopay:=1
	for i:=2;i<=n;i++{
		kopay*=i
	}
	ch <- kopay
}

func main(){
	ch := make(chan int)
	var m int
	fmt.Printf("son kiriting: ")
	fmt.Scanln(&m)
	go factorial(m,ch)
	n := <-ch
	fmt.Println(n)
	
}

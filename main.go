package main

import "fmt"

func main(){
	var num int =1
	list :=LinkedList{data:specialFunc(num),next:nil}
	for num<10{
		num+=2
		list.append(specialFunc(num))
	}
	fmt.Println((list.get(3).(func() int)()))
	list.print()
}

func printSlice(array []int){
	defer fmt.Println()
	for _,v :=range array{
		fmt.Print(" ",v)
	}

}

func specialFunc(num int) func() int{
	return func() int{
		return num*num
	}
}

func testLinkedList(){

}

func testBall(){
	x:= Ball{x:3,y:2,dx:2,dy:2,radius:3}
	x.move()
	x.print()
}

func testInterfaceCastToFunction(dim interface{}){
	val,k:=dim.(func())
	if(k){val()}
}



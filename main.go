package main

import ("fmt";
	bt "./dataStructPract"
	"math/rand"
)

func main(){
	//btreeTest()
	var example string = "9"
	testStringPointer(&example)
	fmt.Println(example)
}

func btreeTest(){
	x:=&bt.Integer{10}
	fmt.Println(x.ToString())

	btreeFirst :=&bt.BTree{}
	btreeFirst.Insert(x,func(){})
	btreeFirst.Insert(&bt.Integer{9},func(){})
	btreeFirst.Insert(&bt.Integer{11},func(){})
	btreeFirst.Print()
	for i:=0;i<13;i++{
		btreeFirst.Insert(&bt.Integer{rand.Intn(100)},func(){})
	}
	btreeFirst.Print()
	v,err := btreeFirst.Smallest().(*bt.Integer) //cast back to Integer pointer which implements comparable
	if(err){fmt.Println(v.ToString())}

}

func testLinkProto(){
	var num int =1
	list :=LinkedList{data:specialFunc(num),next:nil}
	for num<10{
		num+=2
		list.append(specialFunc(num))
	}
	fmt.Println((list.get(3).(func() int)())) // cast to func() int type and execute its value form
	list.remove(4)
	list.remove(4)
	list.print()
	//testInterfaceCastToFunction(func(){fmt.Println("function executed")})
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



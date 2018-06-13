package main

import ("fmt";
	bt "./dataStructPract"
	"math/rand"
)

func main(){
	//btreeTest()
	stackLink()
	//testStackA()
}

func stackLink(){
	stack:=bt.BuildStackL()
	for i:=0;i<10;i++{
		stack.Push(&bt.Integer{i})
	}

	for i:=0;i<10;i++{
		fmt.Println(stack.Pop().(*bt.Integer).ToString())
	}
	fmt.Println(stack.Size())
}

func testStackA(){
	stack:=bt.BuildStack(0)
	for i:=0;i<10;i++{
		stack.Push(&bt.Integer{i})
	}

	for i:=0;i<10;i++{
		fmt.Println(stack.Pop().(*bt.Integer).ToString())
	}
	fmt.Println(stack.ToString())

}

func other(){
	slice1:=[]int{2}
	slice2:=[]int{3}
	testSliceReference(slice1,&slice2)
	fmt.Println(slice1,slice2)
}

func btreeTest(){
	x:=&bt.Integer{10}
	fmt.Println(x.ToString())

	btreeFirst :=&bt.BTree{}//it doesn't have to be a pointer but
	btreeFirst.Insert(x,func(){})
	btreeFirst.Insert(&bt.Integer{9},func(){})//remember comparable interface match signature parameter type which implies *Integer, not Integer
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



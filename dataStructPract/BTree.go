package DataStructPract

import (
	"fmt"
	"strconv"
)

type comparable interface{
	compareTo(other comparable) int
	ToString() string
}

type Integer struct{
	Val int
}

func (this *Integer) compareTo(other comparable) int{
	o,err:=other.(*Integer)
	switch err {
	case this.Val>o.Val:return 1
	case this.Val<o.Val:return -1
	default:return 0
	}

}

func (this *Integer) ToString() string{
	return strconv.Itoa(this.Val)
}

type node struct{
	left, right *node
	Id comparable
	Data interface{}
}

func (this *node) recursiveAdd(new *node){ //test private add for node in package
	if(this.Id.compareTo(new.Id)>0){
		if(this.left==nil){
			this.left = new
			return
		}
		this.left.recursiveAdd(new)
	} else if(this.Id.compareTo(new.Id)<0){
		if(this.right==nil){
			this.right = new
			return
		}
		this.right.recursiveAdd(new)
	} else{
		return //no repeats
	}

}

func (curr *node) toString() string{
	var rightString, leftString string = "",""
	if(curr.right!=nil){
		rightString = curr.right.toString()
	}
	if(curr.left!=nil){
		leftString = curr.left.toString()
	}

	return leftString+" "+curr.Id.ToString()+" "+rightString
}


type BTree struct{
	root *node
	size int

}

func (main *BTree) setRoot(Id comparable,Data interface{})  {
	main.root = &node{Id:Id,Data:Data}
}

func (tree *BTree) Size() int {
	return tree.size
}

//inserts information into BTree
func (main *BTree) Insert(Id comparable,Data interface{}){
	if(main.size==0){
		main.setRoot(Id,Data)
		main.size+=1
		return
	}
	main.root.recursiveAdd(&node{Id:Id,Data:Data})
	main.size+=1
}

func (main *BTree) Search(target comparable) interface{}{
	var internal func (e *node) interface{} // must save variable as type "statically" matching signature of internal method
	internal=func (e *node) interface{}{
		if(e.Id.compareTo(target)>0){
			return internal(e.left)
		}else if(e.Id.compareTo(target)<0){
			return internal(e.right)
		}else if(e.Id.compareTo(target)==0){
			return e.Data
		}else{
			return nil
		}
	}
	return internal(main.root)
}

func (main *BTree) Print(){
	fmt.Println(main.root.toString())
}

func(main *BTree) Smallest() comparable{
	pos:=main.root
	for pos.left!=nil{
		pos = pos.left
	}
	return pos.Id
}

func(main *BTree) Largest() comparable{
	pos:=main.root
	for pos.right!=nil{
		pos = pos.right
	}
	return pos.Id
}







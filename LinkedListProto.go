package main

import (
	"errors"
	"fmt"
)

type LinkedList struct{
	data interface{} // generic for wildtype of any type allowed as data as in int, structs, functions, etc.
	next *LinkedList

}

func (list *LinkedList) append(newest interface{}){
	if(list.next==nil){
		list.next = &LinkedList{data:newest, next:nil} // must be pointer and will check compile time
		return
	}

	list.next.append(newest)
}

func (list* LinkedList) insert(index int, newest interface{}) (err error){
	err = nil // err will be returned at the end of this method with "return" no need to put return err because of signature

	if(index==0){
		list.next = &LinkedList{data:newest, next:nil}
		return
	}
	if(list.next==nil){
		return errors.New("Out of Bounds")
	}
	list.next.insert(index-1,newest)
	return
}

func (list *LinkedList) print(){
	for list!=nil{
		data:=list.data

		v,ok:=data.(func() int) //interface .(type) returns value, boolean error, this statement checks if the data interface{} is func() int type
		if(ok){
			fmt.Println(v())
		} else{
			fmt.Println(v)
		}
		list = list.next

	}

}

func (list *LinkedList) get(index int)(interface{}){
	switch{
		case list==nil:
			return nil
		case index==0:
			return list.data
		default:
			return list.next.get(index-1)
	}
}

func (list *LinkedList) appendList(other *LinkedList){
	for list.next!=nil{
		list = list.next
	}
	list.next = other
}

//removes element at index, not 0 properly
func (list *LinkedList) remove(index int) interface{}{
	for i:=0; i<index-1;i++{
		if(list==nil){
			return errors.New("Out of Bounds")
		}
		list=list.next
	}
	pivot:=list.next
	if pivot==nil{
		return nil
	}
	list.next = pivot.next
	pivot.next=nil
	return pivot.data


}



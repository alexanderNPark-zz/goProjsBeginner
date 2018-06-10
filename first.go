package main

import (
	"fmt"
	"math/rand"
)

func first(){
	fmt.Println(fib(5))
	fmt.Println(fibRecur(5))
	a:=rand.Intn(100)
	var testArray []int = make([]int, a)
	for i:=0; i<a;i++{
		testArray[i] = rand.Intn(10000)
	}

	b :=testArray[:]
	fmt.Println(b)
	selectionSort(b)
	fmt.Println(b)


}

func fib(a uint32) (uint32){
	var(
		i uint32 =1
		temp uint32= 1
		curr uint32= 1
	)

	for i<a {
		curr = curr+temp
		temp = curr-temp
		i+=1
	}
	return curr
}

func bubbleSort(array []int){
	swapped :=true
	decrement:=0
	for swapped{
		swapped = false

		for i:=0;i<len(array)-decrement-1;i++{
			if(array[i]>array[i+1]){
				swapped = true
				temp:=array[i+1]
				array[i+1] = array[i]
				array[i] = temp
			}
		}
		decrement+=1
	}
}

func insertionSort(array []int){

	for index:=1; index<len(array);index++{
			i:=index
			for i>0 && array[i]<array[i-1]{
				temp:=array[i-1]
				array[i-1] = array[i]
				array[i] = temp
				i-=1
			}


	}
}

func selectionSort(array []int){
	for i:=0;i<len(array);i++{
		biggestIndex:=i
		for j:=i;j<len(array);j++{
			if(array[biggestIndex]>array[j]){
				biggestIndex = j
			}
		}
		temp:=array[biggestIndex]
		array[biggestIndex] = array[i]
		array[i] = temp

	}
}

func fibRecur(n int) int{
	if(n<=1){
		return 1
	}
	return fibRecur(n-1)+fibRecur(n-2)
}



package main

import "fmt"

type Ball struct{
	x,y,radius float32
	dx,dy float32
}

func (b *Ball) bounce() bool{
	b.dx = -b.dx
	b.dy = -b.dy
	return true
}

func (main *Ball) compareTo(other *Ball) int{
	if(main.radius>other.radius){
		return 1
	}

	if(main.radius<other.radius){
		return -1
	}

	return 0


}

func (main *Ball) move(){
	main.x+=main.dx
	main.y+=main.dy
}

func (main *Ball) print(){
	fmt.Println("x: ",main.x, "y: ",main.y)
}
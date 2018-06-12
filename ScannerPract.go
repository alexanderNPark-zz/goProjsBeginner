package main

func testStringPointer(content *string){
	*content = "41029"
}

func testSliceReference(slice []int, pointSlice *[]int){
	slice[0] = 0
	(*pointSlice)[0] = 1
}



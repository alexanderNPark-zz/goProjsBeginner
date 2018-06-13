package DataStructPract

type printable interface {
	ToString() string
}

type character struct{
	char string
}

func Character(char string) *character{
	if(len(char)>1){
		return nil
	}
	return &character{char}
}

type singleLinkNode struct{
	data printable
	next *singleLinkNode
}

type stack_L struct{
	head *singleLinkNode
	size int
}

func BuildStackL() *stack_L{
	return &stack_L{}
}

func (stack *stack_L) Size() int{
	return stack.size
}

func (stack *stack_L) Push(data printable){
	stack.head = &singleLinkNode{data,stack.head}
	stack.size+=1
}

func (stack *stack_L) Pop() (data printable){
	if(stack.head==nil){
		return
	}
	data = stack.head.data
	stack.head = stack.head.next
	stack.size-=1
	return
}

func (stack *stack_L) Peek() (data printable){
	data = stack.head.data
	return
}

func (stack *stack_L) ToString() string{
	if(stack.head==nil){
		return "empty"
	}

	var result string
	for temp:=stack.head;temp!=nil;temp = temp.next{
		result = temp.data.ToString()+" "+result
	}
	return result
}



//type of struct using array
type stack_A struct{
	content []printable
	pointer int
}

//abstracted constructor build method for creating stack_A type
func BuildStack(initialSize int) *stack_A{
	if(initialSize==0){
		initialSize=10
	}
	return &stack_A{make([]printable,initialSize),0}

}

//push abstraction for push
func (stack *stack_A) Push(data printable){
	if(stack.pointer==len(stack.content)){
		stack.resize()
	}else{
		stack.content[stack.pointer] = data
		stack.pointer++

	}
}

//returns printable for peek
func (stack *stack_A) Peek() printable{
	if(stack.pointer==0){
		return nil
	}
	return stack.content[stack.pointer-1]
}

//returns the last element LIFO
func (stack *stack_A) Pop() interface{}{
	if(stack.pointer==0){
		return nil
	}
	stack.pointer-=1
	return stack.content[stack.pointer]

}

//private: resizes stack
func (stack *stack_A)resize(){
	new_content:=make([]printable,len(stack.content)*2)
	copy(new_content,stack.content)
	stack.content = new_content
}

// make stack printable by printing stack
func (stack *stack_A) ToString() string{
	if(stack.pointer==0){
		return "empty"
	}
	var result string
	for i:=0;i<stack.pointer;i++{
		result+=" "+stack.content[i].ToString()
	}
	return result
}

func (stack *stack_A) Size() int {
	return stack.pointer
}







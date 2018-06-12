package DataStructPract

type stack_A struct{
	content []interface{}
	pointer int
}

//abstracted constructor build method for creating stack_A type
func BuildStack(initialSize int) *stack_A{
	if(initialSize==0){
		initialSize=10
	}
	return &stack_A{make([]interface{},initialSize),0}

}

//push abstraction for push
func (stack *stack_A) Push(data interface{}){
	if(stack.pointer==len(stack.content)){
		stack.resize()
	}else{
		stack.content[stack.pointer] = data
		stack.pointer++

	}
}

func (stack *stack_A) Pop() interface{}{
	if(stack.pointer==0){
		return nil
	}
	stack.pointer-=1
	return stack.content[stack.pointer]

}

func (stack *stack_A)resize(){
	new_content:=make([]interface{},len(stack.content)*2)
	copy(new_content,stack.content)
	stack.content = new_content
}







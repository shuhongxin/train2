package stack


type Stack struct {
	size int //栈的大小
	top  int //栈顶
	data []int //使用切片创建栈，假使存的数据类型为int型
}

//分配一个新的栈
func CreatStack(size int) Stack {
	newStack := Stack{}
	newStack.size = size
	newStack.data = make([]int,size) 
	return newStack
}

//入栈
func (s Stack) Push(data int) bool {
	if(s.top > s.size){
		return false
	}
	s.data[s.top] = data
	s.top++
	return true
}

//出栈
func (s Stack) Pop() int {
	s.top--
	data := s.data[s.top]
	return data
}
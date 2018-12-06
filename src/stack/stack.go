package stack 

// type
// name
// struct
type Stack struct {
	// push to it
    enque []int
	// tmp queue
    deque []int 
}

// func
// construct
// return
func Constructor() Stack {
	// empty arr, arr 
    return Stack{[]int{}, []int{}}
}

// func
// this
// point to stack
// func param
func (this *Stack) Push(x int)  {
	// append to end
    this.enque = append(this.enque, x)
}

// (this *Stack)
func (this *Stack) Pop() int {
	// len
    length := len(this.enque)
	// skip the last
    for i := 0; i < length-1; i++ {
		// deque starts empty
		// 1st element
        this.deque = append(this.deque, this.enque[0])
		// skip the first
        this.enque = this.enque[1:]
    }

	// enque only has 1 element, so return 
    topEle := this.enque[0]
	// deque replace enque
    this.enque = this.deque
	// nil
    this.deque = nil
   
	// return top element 
    return topEle
}


/** Get the top element. */
func (this *Stack) Top() int {
	// return top element
    topEle := this.Pop()
	// and put it back
    this.enque = append(this.enque, topEle)
    
    return topEle
}


/** Returns whether the stack is empty. */
func (this *Stack) Empty() bool {
    if len(this.enque) == 0 {
        return true
    }
    
    return false
}

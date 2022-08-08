//Design a stack that supports push, pop, top, and retrieving the minimum
//element in constant time.
//
// Implement the MinStack class:
//
//
// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
//
//
// You must implement a solution with O(1) time complexity for each function.
//
//
// Example 1:
//
//
//Input
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//Output
//[null,null,null,null,-3,null,0,-2]
//
//Explanation
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin(); // return -3
//minStack.pop();
//minStack.top();    // return 0
//minStack.getMin(); // return -2
//
//
//
// Constraints:
//
//
// -2³¹ <= val <= 2³¹ - 1
// Methods pop, top and getMin operations will always be called on non-empty
//stacks.
// At most 3 * 10⁴ calls will be made to push, pop, top, and getMin.
//
//
// Related Topics Stack Design 👍 9062 👎 656

package main

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	Main []int
	Min  []int
}

func MinStackConstructor() MinStack {
	stack := MinStack{
		Main: make([]int, 0),
		Min:  make([]int, 0),
	}
	return stack
}

func (this *MinStack) Push(val int) {
	this.Main = append(this.Main, val)
	if len(this.Min) == 0 || this.Min[len(this.Min)-1] >= val {
		this.Min = append(this.Min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.Main) <= 0 {
		return
	}
	val := this.Main[len(this.Main)-1]
	if val == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
	this.Main = this.Main[:len(this.Main)-1]
}

func (this *MinStack) Top() int {
	if len(this.Main) <= 0 {
		return -1
	}
	return this.Main[len(this.Main)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Min) <= 0 {
		return -1
	}
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)

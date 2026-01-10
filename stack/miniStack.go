/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
 

示例 1:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/

/*
思路：首先中等难度的题目，不需要过于复杂的设计，不用太考虑大顶堆小顶堆之类的复杂数据结构
然后常数时间内获取最小值，说明每次push和pop都要维护最小值，这其实也是一个提示。

***另外，一定要分清，是获取还是弹出，如果只是获取，可以通过额外的空间来保存当前的最小值，如果是需要删除就会不一样，就需要设计更复杂的结构。***
***删除会涉及到整个结构的重新组织，顺序的更改，等等，但是获取不会***
*/


type MinStack struct {
    stack  []int
    miniStack []int
}


func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        miniStack: []int{},
    }    
}


func (this *MinStack) Push(val int)  {
    this.stack=append(this.stack,val)
    if len(this.miniStack)==0{
        this.miniStack=append(this.miniStack,val)
    }else{
        cur:=this.miniStack[len(this.miniStack)-1]
        if cur>val{
            this.miniStack=append(this.miniStack,val)
        }else{
            this.miniStack=append(this.miniStack,cur)
        }
    }   
}


func (this *MinStack) Pop()  {
    this.stack=this.stack[:len(this.stack)-1]
    this.miniStack=this.miniStack[:len(this.miniStack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
    
}


func (this *MinStack) GetMin() int {
    return this.miniStack[len(this.miniStack)-1]  

}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
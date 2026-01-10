import "fmt"
/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

 

示例 1：

输入：s = "1 + 1"
输出：2
示例 2：

输入：s = " 2-1 + 2 "
输出：3
示例 3：

输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23
 

提示：

1 <= s.length <= 3 * 105
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
s 表示一个有效的表达式
'+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
'-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
输入中不存在两个连续的操作符
每个数字和运行的计算将适合于一个有符号的 32位 整数
*/
/*
我的calculate思路：
首先，我还是用了那次面试的做法，在里面试图弄多个流程，事实证明这是不可取的，尤其是有空格，很容易出问题
还是把++ 放外层比较好， 循环统一控制，或者有switch 或者 if else 来控制不同的情况
保证里面只走一个情况

我这个的话过了48个case，解决空格问题，括号问题，基本解决了-号作为一元运算的问题，但是有一个case过不了，也不好排查，先这样
实在是找不出bug在哪里了，先放着吧
但是好消息是，实际上面试的时候应该不会有问题。

****所以注意多个操作在for循环里面尽量不要连续，如果不应该是连续的，有很多可能性的话，最好用switch或者if else 来控制不同的情况，保证每次只走一个分支****
除非a走完一定走b，比如说快速排序，要不然后面接各种各样的，一定要及时退出，及时continue
*/

/*注意读题，看能不能简化题目，比如说把括号解开就是很好的例子*/
func calculate_error(s string) int {
    numStack:=[]int{}
    optionStack:=[]int{}  // +1 or -1
    var ans int
    n:=len(s)
    for i:=0;i<n;i++{
        nums:=0
        times:=10
        // for i<n && s[i]>='0' && s[i]<='9'{
        //     nums = nums * times + int(s[i]-'0')
        //     times ++
        //     i++
        // }\
        if  s[i]==' '{
            continue
        }
        if  s[i]>='0' && s[i]<='9'{
            nums = int(s[i]-'0')
            for  i+1<n && s[i+1]>='0' && s[i+1]<='9'{
                nums = nums * times + int(s[i+1]-'0')
                i++
            }
            if len(numStack)==0{
                numStack=append(numStack,nums)  // first number in option
                //fmt.Println("numStack,optionStack",numStack,optionStack)
            }
            if len(numStack)>0 && len(optionStack)>0{
                //fmt.Println("numStack,optionStack before calculate,nums",numStack,optionStack,nums)
                ans =  numStack[len(numStack)-1]+optionStack[len(optionStack)-1]*nums
                numStack[len(numStack)-1]=ans
                optionStack=optionStack[:len(optionStack)-1]
                //fmt.Println("numStack,optionStack after calculate",numStack,optionStack)
            }
            continue
        }                        

        if  s[i]=='+'{
            optionStack=append(optionStack,1)
            //fmt.Println("optionStack append",optionStack)
            continue
        }
        if  s[i]=='-'{
            optionStack=append(optionStack,-1)
            if len(numStack)==0{
                numStack=append(numStack,0)
            }
            //fmt.Println("optionStack append",optionStack)
            continue

        }

        if i<n && s[i]=='('{
            i++
            nums=0
            for i<n && s[i]==' '{
                i++
            }
            if i<n && s[i]=='-'{
                numStack=append(numStack,0)
                optionStack=append(optionStack,-1)
                //fmt.Println("numStack,optionStack (-",numStack,optionStack)
                continue
            }
            if i<n && s[i]>='0' && s[i]<='9'{  // push another number
                nums =  int(s[i]-'0')
                for  i+1<n && s[i+1]>='0' && s[i+1]<='9'{
                    nums = nums * times + int(s[i+1]-'0')
                    i++
                }
                numStack=append(numStack,nums)
            }
            //fmt.Println("numStack,optionStack (",numStack,optionStack)
            continue
        }

        if i<n && s[i]==')' && len(numStack)>1 {                   // pop two number and calculator
            nums1,nums2:=numStack[len(numStack)-2],numStack[len(numStack)-1]
            ans = nums1 + optionStack[len(optionStack)-1]*nums2
            numStack[len(numStack)-2]=ans
            numStack= numStack[:len(numStack)-1]
            optionStack=optionStack[:len(optionStack)-1]
            //fmt.Println("numStack,optionStack )",numStack,optionStack)
            continue
        }

    }
    return numStack[0]    
}

/*
确实只需要符号栈，和一个保存之前算的值的就可以了
用符号栈来存储每个括号里面的运算应该是一个什么样的值，遇到（就存，遇到）就出*/
/*还有就是这是只有+-的，可以考虑直接把括号解开，不需要那么死板*/

func calculate(s string) (ans int) {
    ops := []int{1}
    sign := 1
    n := len(s)
    for i := 0; i < n; {
        switch s[i] {
        case ' ':
            i++
        case '+':
            sign = ops[len(ops)-1]
            i++
        case '-':
            sign = -ops[len(ops)-1]
            i++
        case '(':
            ops = append(ops, sign)
            i++
        case ')':
            ops = ops[:len(ops)-1]
            i++
        default:
            num := 0
            for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
                num = num*10 + int(s[i]-'0')
            }
            ans += sign * num
        }
    }
    return
}

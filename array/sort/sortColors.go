import "fmt"
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]
*/

/*
做题感悟：
- 最好是先在纸上写清楚思路和算法步骤，想清楚节点，怎么初始化，排序的时候什么时候交换，怎么交换
- 为了保证在实际上写题目的过程中不要出错，不要debug太久，像这种++i需要自己控制的时候，用if else 分支更好，防止在中间出现数组越界等情况，同时也保证每次if，else分支都会++，不会死循环
- 注意，在分支中，++ 可能会影响后续操作，尽量往后放，放最后一行
- 另外，如果是不得不在内部用循环++i，比如快速排序，每次实用数组的时候，都要先验证边界条件
*/
func sortColors(nums []int)  {
    firstPosi:=make(map[int]int)
    n:=len(nums)

    // init posi map and i
    i:=0
    firstPosi[nums[i]]=i
    i++
    for i<n {
        if nums[i]>=nums[i-1]{
            if _,ok:=firstPosi[nums[i]];!ok{
                firstPosi[nums[i]]=i
                //fmt.Println("i,nums,posi",i,nums,firstPosi)
            }
            i++
        }else{
            // switch to the smaller front 
            //fmt.Println("posi",firstPosi)
            a,b:=nums[i-1],nums[i]
            nums[firstPosi[a]],nums[i] = b,a
            posi:=firstPosi[a]
            if _,ok:=firstPosi[b];!ok{
                firstPosi[b]=posi
                //fmt.Println("posi update ",a,b,firstPosi,posi)
            }
            firstPosi[a]++

            //fmt.Println("switch 1",i,nums,firstPosi,posi)

            // switch twice ?
            if posi>0 && nums[posi]< nums[posi-1] {
                a,b = nums[posi-1],nums[posi]
                nums[firstPosi[a]],nums[posi] = b,a
                posi=firstPosi[a]
                if _,ok:=firstPosi[b];!ok{
                    firstPosi[b]=posi
                }
                firstPosi[a]++
                //fmt.Println("switch 2",posi,nums,firstPosi)
            }
            i++
        }
    }          
}


/*
Dutch National Flag Algorithm
 更简单的方法：双指针法
 之前也想过这个方法，但是想的还是快排的做法，没想到这么简单

 - 算是快排的的变种，主要的点在于，不是双边遍历，而是，还是从左到右遍历，但是可以双指针，这个可以记录一下
 - 还有，就是如果是3分，那两个指针就可以很自然的得到结果，只要是两个极端就行

 维护两个指针p0和p2，分别表示0和2应该放置的位置
 遍历数组，如果遇到0，就和p0位置交换，并且p0++；
 如果遇到2，就和p2位置交换，并且p2--，注意这里交换完之后，i不++，因为交换过来的元素还需要判断；
 如果遇到1，就直接i++；
 直到i>p2，说明所有2都已经放到最后面了，排序完成


*/
func sortColors1(nums []int) {
    p0, p2 := 0, len(nums)-1
    for i := 0; i <= p2; i++ {
        for ; i <= p2 && nums[i] == 2; p2-- {
            nums[i], nums[p2] = nums[p2], nums[i]
        }
        if nums[i] == 0 {
            nums[i], nums[p0] = nums[p0], nums[i]
            p0++
        }
    }
}
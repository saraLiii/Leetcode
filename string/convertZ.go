//import "fmt"
/*
注意不要重复计算len(s)，放在变量里，不要再循环里面重复计算，包括周期t
*/
func convert(s string, numRows int) string {
    if numRows==1{
        return s
    }
    n:=len(s)
    newS:=make([]byte,0,n)
    for i:=0;i<numRows;i++{
        t:=2*(numRows-1)
        for j:=i;j<n;j+=t{
            if i==0 || i==numRows-1{
                newS=append(newS,s[j])
            }else{
                newS=append(newS,s[j])
                if j+t-2*i <n{
                    newS=append(newS,s[j+t-2*i])
                }     
            }       
        }
        //fmt.Println("numRows,news",numRows,string(newS))
    }
    return string(newS)
}
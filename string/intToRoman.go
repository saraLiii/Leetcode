import "fmt"
/*
稍微有点复杂的方法，先用int转了string再用string去计算的
*/
func intToRoman(num int) string {
    // init num string and times
    var numS []byte
    times:=1
    for num > 0 {
        numS=append(numS,byte(num%10+'0'))
        num=num/10
        times *= 10
    }
    times =times/10  // adjust times
    n:=len(numS)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        numS[i], numS[j] = numS[j], numS[i]
    }
    //fmt.Println("numS times",string(numS),times)
    roman := map[int]byte{
        1:'I',
        5:'V',
        10:'X',
        50:'L',
        100:'C',
        500:'D',
        1000:'M',
    }
    
    var ans []byte
    for _,i:=range numS {
        val:=int(i-'0')
        //fmt.Println("i,val,ans",i,val,string(ans))
        if val==4 || val == 9 {
            ans = append(ans,roman[times],roman[(val+1)*times])
        }else{
            if val>=5 {
                ans = append(ans,roman[5*times])
                val= val-5
            }
            for j:=0;j<val;j++{
                ans = append(ans,roman[times])
            }
        }
        times/=10
    }
    return string(ans)   
}
/*
更好的方法，直接用value去减，从上到下排序就行，我之前也考虑过这个，但是没有想好应该怎么设计数据结构，现在看到了，这种方法确实更简单
*/
var valueSymbols = []struct {
    value  int
    symbol string
}{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func intToRoman1(num int) string {
    roman := []byte{}
    for _, vs := range valueSymbols {
        for num >= vs.value {
            num -= vs.value
            roman = append(roman, vs.symbol...)
        }
        if num == 0 {
            break
        }
    }
    return string(roman)
}

/*
这个是直接给每个位数都硬编码了，实际上面试中可能不适用，但是项目里用会更好
*/

var (
    thousands = []string{"", "M", "MM", "MMM"}
    hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
    return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

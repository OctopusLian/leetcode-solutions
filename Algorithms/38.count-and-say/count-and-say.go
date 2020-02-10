import "strconv"

func countAndSay(n int) string {
    if n<=0{
        return "0"
    }
    return digui("1",n-1)
}

func digui(prestr string,n int)string{
    if n == 0 {
        return prestr
    }
    temp:=""
    for i:=len(prestr);i>0;{
        j:=i
        for ;j>0&&prestr[j-1:j] == prestr[i-1:i];j--{}
        temp=strconv.Itoa(i-j)+prestr[i-1:i]+temp
        i = i-(i-j)
    }
    return digui(temp,n-1)
}
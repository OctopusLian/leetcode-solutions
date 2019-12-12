func reverseString(s []byte)  {
    var result []byte
    for i:=0;i<len(s);i++{
        result = append(result,s[len(s)-i-1])
    }

    for k,v := range result {
        s[k] = v
    }


}

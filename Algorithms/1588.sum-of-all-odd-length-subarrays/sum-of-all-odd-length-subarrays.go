func sumOddLengthSubarrays(arr []int) int {
    length:=len(arr)
    var d int
    var sum int
    var new []int
    for d=1;d<=length;d=d+2{
        for k,_:= range arr{
            if (k+d)<=length{
                new=arr[k:k+d]
                for _,f:=range new{
                    sum=sum+f
                }
            }
        }
    }
    return sum
}
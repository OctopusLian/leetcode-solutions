## 解决  

```go
func countPrimes(n int) int {
    count := 0
    for i:=2;i<n;i++ {
        if (isPrime(i)) {
            count++
        }
    }

    return count
}

func isPrime(n int) bool{
    for i:=2;i<n;i++ {
        if (n%i == 0) {
            //有其它可以被整除的数
            return false
        }
    }
    return true
}
```

超出时间限制。

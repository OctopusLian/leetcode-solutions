/**
 * @param n: An integer
 * @return: An integer
 */
 func dropEggs (n int) int {
    // write your code here
    var ans int
    x := 0
    for ;ans < n; {
        x = x + 1
        ans = ans + x
    }
    
    return x
}
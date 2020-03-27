func hasGroupsSizeX(deck []int) bool {
    deckmap := make(map[int]int)
    if (len(deck) == 0 || len(deck) == 1) {
        return false
    }

    //数组长度为偶数
    for _,num := range deck {
        //记录每个数字出现的次数
        deckmap[num]++
    }
    for _,v1 := range deckmap {
        for _,v2 := range deckmap {
            if(gcd(v1,v2) < 2) {
                return false
            }
        }
    }


    return true
}

func gcd(a,b int) int {
    //求最大公约数
    if (b == 0 ){
        return a
    }
    return gcd(b,a%b)
}

func maxProfit(prices []int) int {
    maxProfit := 0  //最大利润
    for i := 0; i < len(prices)-1; i++ {
        if prices[i+1] > prices[i] {
            //如果后一笔比前一笔数额大，做减法，累积利润
            maxProfit = maxProfit + prices[i+1] - prices[i]
        }
    }
    return maxProfit
}

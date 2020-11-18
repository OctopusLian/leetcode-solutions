func canCompleteCircuit(gas []int, cost []int) int {
	//解法1，一次遍历
    for i, n := 0, len(gas); i < n; {
        sumOfGas, sumOfCost, cnt := 0, 0, 0
        for cnt < n {
            j := (i + cnt) % n
            sumOfGas += gas[j]
            sumOfCost += cost[j]
            if sumOfCost > sumOfGas {
                break
            }
            cnt++
        }
        if cnt == n {
            return i
        } else {
            i += cnt + 1
        }
    }
    return -1
}222222222222222
var d map[string][]string
var ans []string

func findItinerary(tickets [][]string) []string {
	//解法1，深度回溯
    d = map[string][]string{}
    for _, v := range tickets {
        d[v[0]] = append(d[v[0]], v[1])
    }
    for _, v := range d {
        sort.Strings(v)
    }
    ans = []string{}
    dfs("JFK")
    n := len(ans)
	for i := 0; i < n / 2; i++ {
        ans[i], ans[n - i - 1] = ans[n - i - 1], ans[i]
	}
    return ans
}

func dfs(f string) {
    for len(d[f]) > 0 {
        v := d[f][0]
        d[f] = d[f][1: ]
        dfs(v)
    }
    ans = append(ans, f)
}
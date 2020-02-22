func uniquePaths(m int, n int) int {
    dpmap:=make(map[int]map[int]int,0)
    return digui(m-1,n-1,dpmap)
}

func digui(m,n int,dpmap map[int]map[int]int)int{
    if m <= 0 && n <= 0{
        return 1
    }else if nmap,ok := dpmap[m];ok{
        if  v,ok2:= nmap[n];ok2{
            return v
        }
    }
    v1:=0
    v2:=0
    if m>0 {
        v1 = digui(m-1,n,dpmap)
    }
    if n>0{
        v2 = digui(m,n-1,dpmap)
    }

    if _,ok := dpmap[m];!ok{
        dpmap[m] = make(map[int]int,0)
    }
    dpmap[m][n] = v1+v2

    return v1+v2
}

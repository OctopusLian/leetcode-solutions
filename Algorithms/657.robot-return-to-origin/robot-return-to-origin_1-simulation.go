func judgeCircle(moves string) bool {
	//解法1，模拟法
    x, y := 0, 0  //x和y轴的起始值
    for i := 0; i < len(moves); i++ {  //遍历每一个moves的元素
        switch moves[i] {
        case 'U':
            y++
        case 'D':
            y--
        case 'L':
            x--
        case 'R':
            x++
        }
    }
    return x == 0 && y == 0
}
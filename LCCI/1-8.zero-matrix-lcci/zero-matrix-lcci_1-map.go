func setZeroes(matrix [][]int)  {
    //解法1，两个map
    lenth1 := len(matrix)
    lenth2 := len(matrix[0])
    hMap,lMap := make(map[int]bool,10),make(map[int]bool,10)
    for i:=0;i<lenth1;i++{
        for j:=0;j<lenth2;j++{
            if matrix[i][j] == 0{
                hMap[i] = true
                lMap[j] = true
            }
        }
    }
    for i:=0;i<lenth1;i++{
        for j:=0;j<lenth2;j++{
            if hMap[i] || lMap[j]{
                matrix[i][j] = 0
            }
        }
    }
}
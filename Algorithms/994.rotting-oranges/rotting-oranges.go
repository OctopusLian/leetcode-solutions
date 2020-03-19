/*
广度优先搜索，一层一层扩展

思路：与之前的bfs思路不同，初始时，腐烂橘子可能在很多位置上，本题所求就是bfs扩展次数

    将开始时所有腐烂橘子认为一个大腐烂橘子，然后将其沿上下左右扩展
    因为需要计算扩展的层数，所以我们在每一层遍历完成后将层数加1即可
    相比于官方使用map记录时间，通过变量计算一共扩展了几次，空间复杂度更低
*/

func orangesRotting(grid [][]int) int {
    if grid==nil || len(grid)==0{
        return 0
    }
    //按照上右下左方向进行扩展
    dx:=[]int{-1,0,1,0}
    dy:=[]int{0,1,0,-1}
    //行列值
    row:=len(grid)
    col:=len(grid[0])
    res:=0//腐烂完成的时间
    queue:=make([]int,0)
    //首先找到一开始就是腐烂的橘子，将其作为一层
    for i:=0;i<row;i++{
        for j:=0;j<col;j++{
            if grid[i][j]==2{
                //存入映射关系（优秀的方式）
                queue=append(queue,i*col+j)
            }
        }
    }
    //bfs搜索
    for len(queue)!=0{
        res++ //每搜完一层，则时间加一分钟
        cursize:=len(queue)//保存当前层的长度
        for i:=0;i<cursize;i++{
            node:=queue[0]
            queue=queue[1:]
            r,c:=node/col,node%col
            for k:=0;k<4;k++{
                nr:=r+dx[k]
                nc:=c+dy[k]
                if nr>=0 && nr<row && nc>=0 && nc<col && grid[nr][nc]==1{
                    grid[nr][nc]=2  //将新鲜橘子腐烂
                    queue=append(queue,nr*col+nc)//将腐烂橘子入队
                } 
            }
        }
    }
    //判断还有没有新鲜橘子，有就返回-1
    for i:=0;i<row;i++{
        for j:=0;j<col;j++{
            if grid[i][j]==1{
                return -1
            }
        }
    }
    //因为res在计算层的时候，把最开始的腐烂橘子也记为一层，
    //所以结果为res-1
    //存在一个特殊情况，及[[0]]，此时，res就为0，所以不需要-1
    if res==0{
        return res
    }else{
        return res-1
    }
}

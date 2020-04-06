package main

import (
	"strconv"
	"fmt"
)

type Graphmain struct {
	var map [][]int = { 
            {-1, 5, -1, 5, 7}, 
            {-1, -1, 4, -1, -1},
            {-1, -1, -1, 8, 2},
            {-1, -1, 8, -1, 6},
            {-1, 3, -1, -1, -1},
    }
}

func (gm *Graphmain)dfs(end string,path string,maxLength int) {
	if (len(path) - 1 > maxLength) {
		return
	}
	
	if (len(path) > 1 && ath.endsWith(end)) {
		fmt.Println(path + ", "+strconv.Itoa(len(path)))
	}
}



func main() {

}

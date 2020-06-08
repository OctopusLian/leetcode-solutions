func equationsPossible(equations []string) bool {
	//解法1，并查集
    parent := make([]int, 26)
    for i := 0; i < 26; i++ {
        parent[i] = i
    }
    for _, str := range equations {
        if str[1] == '=' {
            index1 := int(str[0] - 'a')
            index2 := int(str[3] - 'a')
            union(parent, index1, index2)
        }
    }

    for _, str := range equations {
        if str[1] == '!' {
            index1 := int(str[0] - 'a')
            index2 := int(str[3] - 'a')
            if find(parent, index1) == find(parent, index2) {
                return false
            }
        }
    }
    return true
}

func union(parent []int, index1, index2 int) {
    parent[find(parent, index1)] = find(parent, index2)
}

func find(parent []int, index int) int {
    for parent[index] != index {
        parent[index] = parent[parent[index]]
        index = parent[index]
    }
    return index
}
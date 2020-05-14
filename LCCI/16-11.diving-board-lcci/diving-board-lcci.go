func divingBoard(shorter int, longer int, k int) []int {
    if k == 0 { // 也就是不需要组合.返回空数组.
		return nil
	}
	if shorter == longer { // 如果长短目录都相等, 最长的目录组就是 k * 长度(shorter, longer)
		return []int{shorter * k}
	}
	group := make([]int, k+1) // 为什么k+1, 因为在组合中我可以使用全是长木板也可以使用全是短木板.
	for i := 0; i < k+1; i++ {
		fmt.Printf("i:%d, k-i:%d\n", i, k-i)
		// 当i=0时,不采用短木板,也就是k-i=k, 表示全使用长木板.
		// 当i=k时,采用全是短木板,而k-k=0,表示不采用长木板
		group[i] = i*shorter + (k-i)*longer
	}
	// 进行排序一下.从低到高的顺序.
	sort.Ints(group)
	return group
}
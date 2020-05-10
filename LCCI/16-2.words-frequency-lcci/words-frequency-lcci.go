type WordsFrequency struct {
    resMap map[string]int  //使用一个map记录一个单词对应的出现频率
}


func Constructor(book []string) WordsFrequency {
    resMap:=make(map[string]int)
	for _,v:=range book{
		resMap[v]++  //记录这本书中每个不重复的单词出现的频率
	}
	return WordsFrequency{resMap}
}


func (this *WordsFrequency) Get(word string) int {
    return this.resMap[word]
}


/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
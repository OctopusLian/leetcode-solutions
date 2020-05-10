func maxAliveYear(birth []int, death []int) int {
    sort.Ints(birth)
    sort.Ints(death)
    Alive := 0
    max := 0
    year := 1900
    j := 0
    for i:=0;i<len(birth);i++ {
        if birth[i] <= death[j] {
            Alive++
            if max < Alive {
                max = Alive
                year = birth[i]
            } 
        } else {
            for birth[i] > death[j] {
                j++
                Alive--
            }
            Alive++
        }
    }
    return year
}